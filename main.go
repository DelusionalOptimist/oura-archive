package main

import (
	"log"

	"github.com/DelusionalOptimist/oura/server"
)

func main() {
	err := server.RunServer()
	if err != nil {
		log.Fatalln("Failed to run API server:", err)
	}
}

