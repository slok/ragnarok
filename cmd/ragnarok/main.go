package main

import (
	"time"

	"github.com/slok/ragnarok/failure"
	"github.com/slok/ragnarok/log"

	_ "github.com/slok/ragnarok/attack/memory"
)

func main() {

	logger := log.Base()
	cfg, _ := failure.ReadConfig([]byte(`
timeout: 30s
attacks:
  - memory_allocation:
      size: 104857600`))

	f, err := failure.NewSystemFailure(cfg, logger)
	if err != nil {
		logger.Fatalf("Error creating system failure: %s", err)
	}

	if err := f.Fail(); err != nil {
		logger.Fatalf("Error Apying  system failure: %s", err)
	}

	time.Sleep(1 * time.Minute)
}
