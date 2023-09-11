package main

import (
	"fmt"
	"log/slog"
	"os"

	"tg-pill-bot/internal/config"

	"github.com/spf13/pflag"
)

func init() { pflag.Parse() }

var (
	configPath = pflag.String("config", "./config.toml", "path to toml config file")
)

func main() {
	config, err := config.ParseTomlFile(*configPath)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse config file: %s", err.Error()))
		os.Exit(1)
	}

	fmt.Println(config)
}
