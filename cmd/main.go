package main

import (
	"github.com/lognitor/go-sdk/log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "test")

	logger.Info("hi there")
	logger.Info("hi there 2")
}

func test() {

}
