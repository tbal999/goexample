package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("Hello Github!")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
