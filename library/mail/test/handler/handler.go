package handler

import (
	"github.com/app-todos/library/logger"
)

// Exec
// return err : エラー情報
func Exec() error {
	logger.Start()
	defer logger.End()

	return nil
}