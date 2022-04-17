package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/onurdogustemel/final-project/route"
)

func main() {
	fmt.Println("Application Starts")
	//Loading Environmental Variable
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Fatal(route.RunAPI(":8090"))

}
