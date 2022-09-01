// テスト配信
// 全体を制御するパッケージ
package handler

import (
	"context"

	"github.com/app-todos/library/logger"
	sg "github.com/app-todos/library/external/sendgrid"
)

const (
	to string = "yusuke040989@gmail.com" // 送信先メールアドレス
)

// Exec
// テスト配信処理の実行
// param ctx : コンテキスト
// return err : エラー情報
func Exec(ctx context.Context) error {
	logger.Log.Info("START")
	defer logger.Log.Info("END")

	// キャンセル処理の設定
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// テスト配信用構造体を設定
	t := sg.NewTest(to)
	if err := t.Send(ctx); err != nil {
		logger.Log.Error(err)
		return err
	}

	return nil
}

// checkCancel
// キャンセルチェック
// param ctx : コンテキスト
// return エラー情報
func CheckCancel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		_ = logger.Log.Error("canceled ctx")
		return ctx.Err()
	default:
		return nil
	}
}