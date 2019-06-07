package main

import (
	"log"
	"os"
)

func main() {
	world := os.Getenv("HELLO")

	if world != "WORLD" {
		log.Fatalf("HELLO env mismatch: %q", world)
	}

	log.Printf("Hello, %s", world)
}
