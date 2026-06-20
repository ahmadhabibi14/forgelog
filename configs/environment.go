package configs

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvDev  string = `dev`
	EnvStag string = `stag`
	EnvProd string = `prod`
)

func GetProjectEnv() string {
	return os.Getenv(`PROJECT_ENV`)
}

func LoadEnv() {
	dirRetryList := []string{``, `../`, `../../`, `../../../`}
	var err error
	for _, dirPrefix := range dirRetryList {
		envFile := dirPrefix + `.env`

		err = godotenv.Overload(envFile + `.dev`)
		if err == nil {
			slog.Info(`file .env.dev loaded (development environment)`)
			return
		}

		err = godotenv.Overload(envFile)
		if err == nil {
			slog.Info(`file .env loaded (staging/production environment)`)
			return
		}
	}
	panic("cannot load .env file")
}
