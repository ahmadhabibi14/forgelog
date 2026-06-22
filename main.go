package main

import (
	"context"
	"forgelog/configs"
	cronjob "forgelog/internal/cron"
	"forgelog/internal/handler"
	"forgelog/internal/lib/docker"
	"forgelog/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
	"golang.org/x/sync/errgroup"
)

func main() {
	configs.LoadEnv()
	logger.InitializeLogger()

	dockerClient, err := docker.InitDockerClient()
	if err != nil {
		logger.Log.Fatal(err.Error(), "failed to initialize docker client")
		os.Exit(1)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "PATCH",
		},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	app.Get("/*", static.New("./views/dist"))
	app.Get("/", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Status(fiber.StatusOK).SendFile("./views/dist/index.html")
	})

	app.Get("/healthz", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.Status(fiber.StatusOK).SendString("ok")
	})

	handler.NewHandlerContainer(app, dockerClient)
	handler.NewHandlerSystem(app)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	group, groupCtx := errgroup.WithContext(ctx)

	// Start HTTP Server
	group.Go(func() error {
		err := app.Listen(":3000")
		return err
	})

	// Start Cronjob
	group.Go(func() error {
		cronjob.CollectStats(groupCtx)
		return nil
	})

	// Close gracefully
	group.Go(func() error {
		<-groupCtx.Done()
		dockerClient.Close()
		return app.Shutdown()
	})

	err = group.Wait()
	switch err {
	case context.Canceled, nil:
		logger.Log.Info("server stopped")
	default:
		logger.Log.Error(err, "failed to stop server")
	}
}
