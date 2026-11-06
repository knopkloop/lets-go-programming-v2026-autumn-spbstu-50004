package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

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

	if err := saveJSON(cfg.OutputFile, rates); err != nil {
		panic(fmt.Errorf("save result: %w", err))
	}
}

func saveJSON(path string, rates []cbr.Currency) error {
	data, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}
