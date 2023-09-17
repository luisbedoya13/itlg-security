package main

import (
	"github.com/subosito/gotenv"
	"log"
)

func init() {
	if err := gotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading variables: %v\n", err)
	}
}

func main() {}
