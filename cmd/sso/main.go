package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Print(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложения 
	
	// TODO: запустить gRPC-сервер приложения
}

