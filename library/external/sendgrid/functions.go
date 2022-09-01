// SendGridのAPI処理用パッケージ
package sendgrid

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/app-todos/library/logger"
)

// CreateBatchIDのレスポンス格納用構造体
type batchIDRes struct {
	BatchID string `json:"batch_id"`
}

// CreateBatchID
// バッチID生成APIを実行
// param ctx : コンテキスト
// return バッチID
// return エラー情報
func (c *Client) CreateBatchID(ctx context.Context) (string, error) {
	// API実行
	res, err := c.Post(ctx, "/mail/batch", nil)
	if err != nil {
		return "", err
	}
	// ステータスコード201以外の場合
	if res.StatusCode != 201 {
		return "", errors.New(fmt.Sprintf("バッチID生成に失敗しました/status:%d/body: %s", res.StatusCode, res.Body))
	}

	// 生成したバッチIDを格納
	var batchIDRes batchIDRes
	b := []byte(res.Body)
	if err := json.Unmarshal(b, &batchIDRes); err != nil {
		logger.Log.Debug("JSONのパースに失敗しました")
		return "", err
	}

	return batchIDRes.BatchID, nil
}

// ValidateBatchID
// バッチIDの有効チェック
// param ctx : コンテキスト
// param batchID : バッチID
// return エラー情報
func (c *Client) ValicateBatchID(ctx context.Context, batchID string) error {
	uri := fmt.Sprintf("/mail/batch/%s", batchID)
	// API実行
	res, err := c.Get(ctx, uri)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		return nil
	case 400:
		logger.Log.Warnf("batchIDが無効です/batchID: %s", batchID)
		return nil
	default:
		return errors.New(fmt.Sprintf("%d: %s/batchID: %s", res.StatusCode, res.Body, batchID))
	}
}

// Send
// メール送信
// param ctx : コンテキスト
// param req : リクエストボディ
// return エラー情報
func (c *Client) Send(ctx context.Context, reqBody []byte) error {
	res, err := c.Post(ctx, "/mail/send", reqBody)
	if err != nil {
		return err
	}
	// ステータスコード202以外の場合
	if res.StatusCode != 202 {
		return errors.New(fmt.Sprintf("メール送信に失敗しました/status:%d/body: %s", res.StatusCode, res.Body))
	}

	return nil
}