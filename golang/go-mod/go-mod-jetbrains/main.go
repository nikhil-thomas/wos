package main

import "fmt"
import "github.com/pkg/errors"

func main() {
	err := HelloWorld("Hello World")
	if err != nil {
		panic(err)
	}
}

func HelloWorld(message string) error {
	_, err := fmt.Println(message)
	return err
}
