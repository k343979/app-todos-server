// SendGridAPI通信用パッケージ
package sendgrid

import (
	"context"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// API通信用クライアント
type Client struct {
	Host   string // https://api.sendgrid.com/v3
	ApiKei string // APIキー
}

// API通信用クライアント構造体を生成
// return : *Client
func NewClient() *Client {
	return &Client{
		Host:   os.Getenv("SENDGRID_API_URL"),
		ApiKei: os.Getenv("SENDGRID_API_KEY"),
	}
}

// Get
// API実行メソッド(GET)
// param ctx : コンテキスト
// param uri : APIエンドポイント
// return レスポンス結果
// return エラー情報
func (c *Client) Get(ctx context.Context, uri string) (*rest.Response, error) {
	// リクエスト情報のセット
	req := sendgrid.GetRequest(c.ApiKei, uri, c.Host)
	req.Method = rest.Method(rest.Get)
	// APIの実行
	return c.API(ctx, req)
}

// Post
// API実行メソッド(POST)
// param ctx : コンテキスト
// param uri : APIエンドポイント
// param req : リクエストボディ
// return レスポンス結果
// return エラー情報
func (c *Client) Post(ctx context.Context, uri string, reqBody []byte) (*rest.Response, error) {
	// リクエスト情報のセット
	req := sendgrid.GetRequest(c.ApiKei, uri, c.Host)
	req.Method = rest.Method(rest.Post)
	req.Body = reqBody
	// APIの実行
	return c.API(ctx, req)
}

// API
// SendGridにリクエストを送信
// param ctx : コンテキスト
// param req : リクエスト
// return レスポンス結果
// return エラー情報
func (c *Client) API(ctx context.Context, req rest.Request) (*rest.Response, error) {
	return sendgrid.MakeRequestWithContext(ctx, req)
}