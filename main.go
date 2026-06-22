package main

import (
	"context"
	"forgelog/configs"
	cronjob "forgelog/internal/cron"
	"forgelog/internal/handler"
	"forgelog/internal/lib/logger"
	"io"
	"log"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/moby/moby/client"
)

var (
	dockerClient *client.Client
)

func init() {
	cli, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	dockerClient = cli

	configs.LoadEnv()
	logger.InitializeLogger()
}

func dockerTerminal(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		logger.Log.Info("Connection closed")
	}()
	containerId := conn.Params("container_id")
	ctx := context.Background()

	execResp, err := dockerClient.ExecCreate(ctx, containerId, client.ExecCreateOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		TTY:          true,
		Cmd: []string{
			"/bin/bash",
		},
	})

	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	attachResp, err := dockerClient.ExecAttach(
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

func listContainersDocker(c fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	containers, err := dockerClient.ContainerList(c.Context(), client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(containers)
}

func getContainerDetail(c fiber.Ctx) error {
	containerId := c.Params("container_id")

	res, err := dockerClient.ContainerInspect(context.Background(), containerId, client.ContainerInspectOptions{
		Size: true,
	})
	if err != nil {
		logger.Log.Error(err, "failed to get container info")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func main() {
	go cronjob.CollectStats()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "PATCH",
		},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	app.Get("/api/system/stats", handler.GetStats)
	app.Get("/api/system/terminal", websocket.New(handler.TerminalHandler))
	app.Post("/api/containers/docker", listContainersDocker)
	app.Get("/api/containers/docker/terminal/:container_id", websocket.New(dockerTerminal))
	app.Get("/api/containers/docker/detail/:container_id", getContainerDetail)
	app.Post("/api/containers/docker/detail/:container_id", getContainerDetail)
	app.Post("/api/containers/docker/config/:container_id", getContainerDetail)

	app.Get("/healthz", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.Status(fiber.StatusOK).SendString("ok")
	})

	log.Fatal(app.Listen(":3000"))
}
