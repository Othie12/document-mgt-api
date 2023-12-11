package main

import (
	"hrms/routes"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file: " + err.Error())
	}
}

func main() {

	routes.Routes()
}
