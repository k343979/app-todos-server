package logger

import (
	"log"
)

var process *string

// 初期設定
func Set(processName string) {
	process = &processName
}

// 処理開始ログ
func Start() {
	log.Printf("[START]: %s\n", *process)
}

// 処理終了ログ
func End() {
	log.Printf("[END]: %s\n", *process)
}