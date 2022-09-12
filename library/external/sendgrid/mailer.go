// メール組み立て用パッケージ
package sendgrid

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/app-todos/library/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// 送信対象情報
type Target struct {
	Name  string
	Email string
}

// メール基本情報
type Info struct {
	Target    *Target // 送信対象情報
	PathHtml  string  // HTMLメール本文のテンプレートパス
	PathPlain string  // テキストメール本文のテンプレートパス
}

// テスト配信用構造体
type Test struct {
	Client *Client // API通信用クライアント
	Info   *Info   // メール基本情報
}

// メールインタ-フェース
type Mail interface {
	Send(context.Context) error // メール送信処理
}

// NewTarget
// Target構造体を生成
// param name : 送信対象者名
// param email : 送信先メールアドレス
// return *Target
func NewTarget(name, email string) *Target {
	return &Target{
		Name:  name,
		Email: email,
	}
}

// NewTest
// Test構造体をMailインターフェースで生成
// return Mailインターフェース
func (t *Target) NewTest(html, plain string) Mail {
	return &Test{
		Client: NewClient(),
		Info:   &Info{
			Target:    t,
			PathHtml:  html,
			PathPlain: plain,
		},
	}
}

// (t *Test) Send
// テスト配信処理
// param ctx : コンテキスト
// return エラー情報
func (t *Test) Send(ctx context.Context) error {
	// API通信用クライアント
	c, info := t.Client, t.Info
	// バッチIDを生成
	batchID, err := c.CreateBatchID(ctx)
	if err != nil {
		return err
	}

	// batchIDの有効チェック
	if err := c.ValidateBatchID(ctx, batchID); err != nil {
		err = logger.Log.Errorf("err: %w", err)
		return err
	}

	// メール情報を組立
	reqBody := info.Build(batchID)

	// メール送信
	return c.Send(ctx, reqBody)
}

// Build
// メール情報の組立
// param batchID : バッチID
// return メールのリクエスト内容
func (info *Info) Build(batchID string) []byte {
	// メール基本情報
	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("【水口テスト】事務局", os.Getenv("SENDGRID_FROM")))
	m.Subject = "テストメール"

	// 送信対象をセット
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(info.Target.Name, info.Target.Email))
	// 送信対象者用に置換文字列をセット
	info.SetConv(p)
	// 送信対象をメール情報にセット
	m.AddPersonalizations(p)

	// Content-Typeの設定(HTML形式を優先)
	contentType := "text/plain"
	// ファイルパスの設定
	fp := info.PathPlain
	if info.PathHtml != "" {
		contentType = "text/html"
		fp = info.PathHtml
	}

	// テンプレートファイルの読み込み
	buf, err := ioutil.ReadFile(fp)
	if err != nil {
		_ = logger.Log.Errorf("テンプレートファイルの読み込みに失敗しました/%w", err)
	}

	// メール内容のセット
	content := mail.NewContent(contentType, string(buf))
	m.AddContent(content)

	// バッチIDをセット
	m.SetBatchID(batchID)

	return mail.GetRequestBody(m)
}