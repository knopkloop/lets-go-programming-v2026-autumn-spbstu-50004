package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/knopkloop/task-3/internal/cbr"
	"github.com/knopkloop/task-3/internal/config"
)

func main() {
	configPath := flag.String("config", "", "path to YAML configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	rates, err := cbr.Load(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("load rates: %w", err))
	}

	result, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		panic(fmt.Errorf("marshal json: %w", err))
	}

	fmt.Println(string(result))
}
