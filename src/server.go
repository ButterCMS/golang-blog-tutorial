package main

import (
	"os"

	"github.otrex/blog/src/config"
	"github.otrex/blog/src/router"
)

func main() {
	config.SetupConfig("../.env")
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
