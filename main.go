package main

import (
	"context"
	"io"
	"log"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/moby/moby/client"
)

func terminalHandler(conn *websocket.Conn) {
	defer conn.Close()

	containerId := conn.Params("container_id")

	ctx := context.Background()

	cli, err := client.New()
	if err != nil {
		_ = conn.WriteMessage(
			websocket.TextMessage,
			[]byte(err.Error()),
		)
		return
	}

	execResp, err := cli.ExecCreate(ctx, containerId, client.ExecCreateOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		TTY:          true,
		Cmd: []string{
			"/bin/bash",
		},
	})

	if err != nil {
		conn.WriteMessage(
			websocket.TextMessage,
			[]byte(err.Error()),
		)
		return
	}

	attachResp, err := cli.ExecAttach(
		ctx,
		execResp.ID,
		client.ExecAttachOptions{
			TTY: true,
		},
	)
	if err != nil {
		conn.WriteMessage(
			websocket.TextMessage,
			[]byte(err.Error()),
		)
		return
	}
	defer attachResp.Close()

	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := attachResp.Reader.Read(buf)

			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}
				return
			}

			err = conn.WriteMessage(
				websocket.TextMessage,
				buf[:n],
			)
			if err != nil {
				return
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		_, err = attachResp.Conn.Write(msg)
		if err != nil {
			return
		}
	}
}

func main() {
	app := fiber.New()

	app.Use("/ws", func(c fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:container_id", websocket.New(terminalHandler))

	app.Get("/", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendFile("./views/index.html")
	})

	log.Fatal(app.Listen(":3000"))
}
