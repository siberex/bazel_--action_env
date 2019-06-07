package main

import (
	"log"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	world := os.Getenv("HELLO")

	if world != "WORLD" {
		t.Errorf("HELLO env mismatch: %q", world)
	}

	log.Printf("Hello, %s", world)
}
