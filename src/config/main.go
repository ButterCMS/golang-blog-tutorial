package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.otrex/blog/src/utils"
)

func SetupConfig(path string) {
	err := godotenv.Load(utils.AbsPath(path))
	if err != nil {
		panic("no .env file found")
	}
	os.Setenv("TEMPLATE_PATH", utils.AbsPath("templates"))
	os.Setenv("ASSETS_PATH", utils.AbsPath("assets"))
	os.Setenv("BUTTERCMS_BASE_URL", "https://api.buttercms.com/v2/")
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8000")
	}
}
