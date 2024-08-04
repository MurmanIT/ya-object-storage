package main

import (
	"fmt"
	"ya-storage/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("config error: %s", err)
		return
	}
	fmt.Println(cfg)
}

// TODO: init logger
// TODO: init server
// TODO: init storage
// TODO: run server
