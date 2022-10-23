package main

import (
	"flag"
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
	filePath := flag.String("cfg", "", "path to config.yaml")
	flag.Parse()

	if filePath != nil {
		if *filePath == "" {
			fmt.Println("Please, provide path to config.yaml")
		}

		s, err := server.NewServer(*filePath)
		if err != nil {
			fmt.Printf("INIT ERROR: %s", err.Error())
			return
		}

		fmt.Printf("%+v\n", s.Cfg.Auth.HeaderKey)

		if err := s.Run(); err != nil {
			fmt.Printf("RUN ERROR: %s", err.Error())
			return
		}
	} else {
		fmt.Println("NO CONFIG PATH PROVIDED")
	}

}
