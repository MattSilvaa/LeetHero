package config

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Delay           time.Duration
	Headless        bool
	LeetCodeSession string
	Problems        []string
}

func Load() *Config {
	cfg := &Config{}

	// Command line flags
	flag.StringVar(&cfg.LeetCodeSession, "cookie", "", "LeetCode session cookie")
	flag.BoolVar(&cfg.Headless, "headless", true, "Run in headless mode")
	flag.DurationVar(&cfg.Delay, "delay", 5*time.Second, "Delay between actions")
	problems := flag.String("problems", "two-sum,add-two-numbers", "Comma-separated problem slugs")

	flag.Parse()

	if cfg.LeetCodeSession == "" {
		cfg.LeetCodeSession = os.Getenv("LEETCODE_SESSION")
	}

	if cfg.LeetCodeSession == "" {
		log.Fatal("LeetCode session cookie is required. Set via -cookie flag or LEETCODE_SESSION environment variable")
	}

	// Parse problems list
	cfg.Problems = strings.Split(*problems, ",")

	return cfg
}
