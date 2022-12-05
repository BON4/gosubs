package main

import (
	"fmt"

	"github.com/BON4/gosubs/internal/server"
)

// @title           Telegram Subs API
// @version         1.0
// @description     This service provide functionality for storing and managing privat telegram channels with subscription based payments for acessing content.

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apiKey JWT
// @in header
// @name authorization
func main() {
	s, err := server.NewServer(".")
	if err != nil {
		fmt.Printf("INIT ERROR: %s", err.Error())
		return
	}

	fmt.Printf("%+v\n", s.Cfg.HeaderKey)

	if err := s.Run(); err != nil {
		fmt.Printf("RUN ERROR: %s", err.Error())
		return
	}
}
