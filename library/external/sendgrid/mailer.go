// メール組み立て用パッケージ
package sendgrid

import (
	"context"
	"os"

	"github.com/app-todos/library/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// 送信対象
type Target struct {
	From string // 送信元メールアドレス
	To   string // 送信先メールアドレス
}

// テスト配信用構造体
type Test struct {
	Client *Client // API通信用クライアント
	Target *Target // 送信対象
}

// メールインタ-フェース
type Mail interface {
	Send(context.Context) error // メール送信処理
}

// NewTest
// Test構造体をMailインターフェースで生成
// param to : 送信先メールアドレス
// return Mailインターフェース
func NewTest(to string) Mail {
	return &Test{
		Client: NewClient(),
		Target: &Target{
			From: os.Getenv("SENDGRID_FROM"),
			To:   to,
		},
	}
}

// (t *Test) Send
// テスト配信処理
// param ctx : コンテキスト
// return エラー情報
func (t *Test) Send(ctx context.Context) error {
	// API通信用クライアント
	c, target := t.Client, t.Target
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// batchIDの有効チェック
	if err := c.ValicateBatchID(ctx, batchID); err != nil {
		err = logger.Log.Errorf("err: %w", err)
		return err
	}

	// メール基本情報
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("yusuke", target.From))
	m.Subject = "テスト送信"

	// 送信対象をセット
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail("yusuke様", target.To))
	// 送信対象をメール情報にセット
	m.AddPersonalizations(p)

	// メール本文をセット
	content := mail.NewContent("text/html", "<br /><p>テスト送信だよん</p>")
	m.AddContent(content)

	// バッチIDをセット
	m.SetBatchID(batchID)

	// メール情報をリクエスト用にセット
	reqBody := mail.GetRequestBody(m)

	// メール送信
	return c.Send(ctx, reqBody)
}