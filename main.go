package main

import (
	"os"

	"quest/route"
)

func main() {
	route := route.Register()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	route.Run(":" + port)
}
