// テスト配信
// 全体を制御するパッケージ
package testmail

import (
	"context"

	"github.com/app-todos/library/logger"
	sg "github.com/app-todos/library/external/sendgrid"
)

const (
	name  string = "水口佑介"
	email string = "yusuke040989@gmail.com" // 送信先メールアドレス
)

const (
	html string  = "/go/src/github.com/app-todos/library/mail/template/testmail/html/text.html"
	plain string = "/go/src/github.com/app-todos/library/mail/template/testmail/plain/text.html"
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

	// 送信対象を設定
	target := sg.NewTarget(name, email)
	// テスト配信用構造体を設定
	t := target.NewTest(html, plain)
	// メールを送信
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