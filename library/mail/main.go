package main

import (
	"github.com/app-todos/library/logger"
)

func init() {
	logger.Set("testmail")
}

func main() {
	logger.Log.Info("sr")
}