package main

import (
	"grpc-todolist/user/config"
	"grpc-todolist/user/internal/repository"
)

func main() {
	config.InitConfig()
	repository.InitDB()
}
