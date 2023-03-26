package authapp

import (
	"os"

	"github.com/Sonu875/go-auth/logger"
)

func santityCheck() {
	switch {
	case os.Getenv("APP_HOST") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("APP_PORT") == "":
		logger.Fatal("Application port is not set as evn")
	case os.Getenv("DB_HOST") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_PORT") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_PASSWORD") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_NAME") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_USER") == "":
		logger.Fatal("Application host is not set as evn")
	}
}

func Start() {

}
