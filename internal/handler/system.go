package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"forgelog/internal/lib/logger"
	"forgelog/internal/state"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

type StatsResponse struct {
	Memory    string `json:"memory"`
	Disk      string `json:"disk"`
	CPU       string `json:"cpu"`
	Bandwidth string `json:"bandwidth"`
}

func GetStats(c fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("X-Accel-Buffering", "no")

	c.SendStreamWriter(func(w *bufio.Writer) {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-c.Context().Done():
				logger.Log.Info("Client disconnected")
				return
			case <-ticker.C:
				state.SystemStats.Mu.RLock()
				data, err := json.Marshal(state.SystemStats.Stat)
				state.SystemStats.Mu.RUnlock()

				if err != nil {
					return
				}

				_, err = fmt.Fprintf(w, "event: stats\n")
				if err != nil {
					return
				}

				_, err = fmt.Fprintf(w, "data: %s\n\n", data)
				if err != nil {
					return
				}

				if err := w.Flush(); err != nil {
					return
				}
			}
		}
	})

	return nil
}

func TerminalHandler(c *websocket.Conn) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Log.Error(err)
		return
	}

	var shell string = "/bin/sh"
	if _, err := os.Stat("/bin/bash"); err == nil {
		shell = "/bin/bash"
	}

	cmd := exec.Command(shell)
	cmd.Dir = homeDir

	ptmx, err := pty.Start(cmd)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	defer func() {
		ptmx.Close()
		cmd.Process.Kill()
	}()

	// PTY -> WebSocket
	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				if err != io.EOF {
					logger.Log.Error(err)
				}
				_ = c.Close()
				return
			}

			if err := c.WriteMessage(websocket.TextMessage, buf[:n]); err != nil {
				return
			}
		}
	}()

	// WebSocket -> PTY
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			break
		}

		if _, err := ptmx.Write(msg); err != nil {
			break
		}
	}
}

func humanizeBytes(b uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case b >= TB:
		return fmt.Sprintf("%.1fTB", float64(b)/float64(TB))
	case b >= GB:
		return fmt.Sprintf("%.1fGB", float64(b)/float64(GB))
	case b >= MB:
		return fmt.Sprintf("%.1fMB", float64(b)/float64(MB))
	case b >= KB:
		return fmt.Sprintf("%.1fKB", float64(b)/float64(KB))
	default:
		return fmt.Sprintf("%dB", b)
	}
}
