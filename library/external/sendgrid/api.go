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

// SendGridインターフェース
type SendGrid interface {
	ExecApi(uri, method string) (*rest.Response, error)
}

// API通信用クライアント構造体をSendGridインターフェースとして返却
// return : *Client SendGridインターフェース
func New() SendGrid {
	return &Client{
		Host:   os.Getenv("SENDGRID_API_URL"),
		ApiKei: os.Getenv("SENDGRID_API_KEY"),
	}
}

// ExecApi
// API実行メソッド
// param uri : APIエンドポイント
// param method : API通信種別
// return res : レスポンス結果
// return err : エラー情報
func (c *Client) ExecApi(uri, method string) (res *rest.Response, err error) {
	// リクエスト情報のセット
	req := sendgrid.GetRequest(c.ApiKei, uri, c.Host)
	req.Method = rest.Method(method)
	// APIの実行
	res, err = sendgrid.API(req)
	return
}