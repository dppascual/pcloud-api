package main

import (
	"fmt"
	"os"
)

func main() {
	api := API{}
	api.Init()

	// Read the environment variable API_PORT
	apiPort := os.Getenv("API_PORT")

	if apiPort == "" {
		apiPort = "80"
	}

	api.Run(fmt.Sprintf(":%s", apiPort))
}
