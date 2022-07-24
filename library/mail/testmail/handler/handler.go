package handler

import (
	"fmt"

	"github.com/app-todos/library/logger"
	sg "github.com/app-todos/library/external/sendgrid"
)

// Exec
// return err : エラー情報
func Exec() error {
	// ログ設定
	logger.Set("mail/testmail")
	logger.Start()
	defer logger.End()

	// (仮) batchID
	batchID := "test"

	// batchIDの有効チェック
	res, err := sg.ValicateBatchID(batchID)
	if err != nil {
		return err
	}

	// ステータスコードが200以外の場合、エラーにして返却
	if res.StatusCode != 200 {
		err = fmt.Errorf("batchIDが無効です/batchID: %s", batchID)
	}

	return nil
}