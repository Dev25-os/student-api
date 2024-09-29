package main

import (
	"fmt"

	"github.com/Dev25-os/student-api/internal/config"
)

func main() {
	fmt.Println("Welcome to student api")
	//load config
	cfg := config.MustLoad()
	fmt.Print("Config file", cfg)
	// database setup
	//setup router
	//setup server

}
