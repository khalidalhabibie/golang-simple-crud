package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"golang-simple-crud/route"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Main Application Starts")
	//Loading Environmental Variable
	loadenv()

	log.Fatal(route.Run(":8081"))

}
