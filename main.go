package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/khalidalhabibie/depatu/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Starts")
	log.Fatal(route.Run(":8081"))

}
