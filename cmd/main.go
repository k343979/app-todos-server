// mainパッケージ
// サーバ起動
package main

import (
	"context"

	"github.com/app-todos/cmd/infrastructure/router"
	"github.com/app-todos/library/logger"
)

// 初期化関数
func init() {
	logger.Set("todos")
}

// エントリポイント
func main() {
	ctx := context.Background()
	router.Run(ctx)
}
