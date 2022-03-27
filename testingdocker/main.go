package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Init func launched")

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	truth, ok := os.LookupEnv("AINUR")
	if ok {
		fmt.Println("Truth:", truth)
	}
}
