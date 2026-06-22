package handler

import (
	"context"
	"forgelog/internal/lib/logger"
	"io"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/moby/moby/client"
)

type handlerContainer struct {
	dockerClient *client.Client
}

func NewHandlerContainer(app *fiber.App, dockerClient *client.Client) {
	h := &handlerContainer{
		dockerClient: dockerClient,
	}

	app.Post("/api/containers", h.List)

	app.Get("/api/container/:id", h.GetById)
	app.Get("/api/container/:id/terminal", websocket.New(h.Terminal))
	app.Post("/api/container/:id/config", h.GetById)
}

func (h *handlerContainer) Terminal(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		logger.Log.Info("Connection closed")
	}()
	containerId := conn.Params("id")
	ctx := context.Background()

	execResp, err := h.dockerClient.ExecCreate(ctx, containerId, client.ExecCreateOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		TTY:          true,
		Cmd:          []string{"/bin/bash"},
	})

	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	attachResp, err := h.dockerClient.ExecAttach(
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
					logger.Log.Error(err)
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

func (h *handlerContainer) List(c fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	containers, err := h.dockerClient.ContainerList(c.Context(), client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(containers)
}

func (h *handlerContainer) GetById(c fiber.Ctx) error {
	containerId := c.Params("id")

	res, err := h.dockerClient.ContainerInspect(context.Background(), containerId, client.ContainerInspectOptions{
		Size: true,
	})
	if err != nil {
		logger.Log.Error(err, "failed to get container info")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
