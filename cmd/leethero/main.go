package main

import (
	"log"

	"github.com/MattSilvaa/leethero/internal/bot"
	"github.com/MattSilvaa/leethero/internal/config"
)

func main() {
	cfg := config.Load()

	hero, err := bot.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create LeetHero: %v", err)
	}

	err = hero.Run()
	if err != nil {
		log.Fatalf("Failed to run LeetHero: %v", err)
	}

	log.Print("Solved all your problems 🦸...you can thank me later. Goodbye!")
}
