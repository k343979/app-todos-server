// SendGridAPI用パッケージ
package sendgrid

import (
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// SendgGridのAPI通信用クライアント
type Client struct {
	Host   string // https://api.sendgrid.com/v3
	ApiKei string // APIキー
}

// API通信用クライアント構造体を返却
// return : *Client
func New() *Client {
	return &Client{
		Host:   os.Getenv("SENDGRID_API_URL"),
		ApiKei: os.Getenv("SENDGRID_API_KEY"),
	}
}

// Get
// API実行メソッド(GET)
// param uri : APIエンドポイント
// return res : レスポンス結果
// return err : エラー情報
func (c *Client) Get(uri string) (*rest.Response, error) {
	// リクエスト情報のセット
	req := sendgrid.GetRequest(c.ApiKei, uri, c.Host)
	req.Method = rest.Method("GET")
	// APIの実行
	return sendgrid.API(req)
}

// Post
// API実行メソッド(POST)
// param uri : APIエンドポイント
// return res : レスポンス結果
// return err : エラー情報
func (c *Client) Post(uri string, reqBody []byte) (*rest.Response, error) {
	// リクエスト情報のセット
	req := sendgrid.GetRequest(c.ApiKei, uri, c.Host)
	req.Method = rest.Method("POST")
	if reqBody != nil {
		req.Body = reqBody
	}
	// APIの実行
	return sendgrid.API(req)
}
