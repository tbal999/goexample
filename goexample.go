package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Github!")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
