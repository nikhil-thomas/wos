package main

import (
	"io"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		logrus.Fatal(err)
	}

}
