package sendgrid

import (
	"fmt"

	"github.com/sendgrid/rest"
)

// ValidateBatchID
// batchIDの有効チェック
// param batchID : バッチID
// return レスポンス結果
// return エラー情報
func ValicateBatchID(batchID string) (*rest.Response, error) {
	// SendGridAPIのセット
	sgClient := New()
	uri := fmt.Sprintf("/mail/batch/%s", batchID)
	// API実行
	return sgClient.Get(uri)
}