package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)

	data := User{
		Name:    "Ainur",
		Country: "KZ",
	}

	if err := enc.Encode(data); err != nil {
		log.Fatal(err)
	}

	file.Write(buf.Bytes())
	bslice, _ := os.ReadFile(file.Name())
	fmt.Println(string(bslice))
}
