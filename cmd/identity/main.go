package main

import (
	"fmt"
	"identity/internal/config"
)

func main() {
	//инициализровать конфиг
	cfg := config.MustLoad()

	fmt.Println(cfg)
	//инициализровать логгер

	//инициализровать приложение

	//запустить grpc сервер
}
