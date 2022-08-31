package handler

import (
	"github.com/app-todos/library/logger"
	sg "github.com/app-todos/library/external/sendgrid"
)

// Exec
// return err : エラー情報
func Exec() error {
	logger.Log.Info("START")
	defer logger.Log.Info("END")

	// (仮) batchID
	batchID := "test"

	// batchIDの有効チェック
	res, err := sg.ValicateBatchID(batchID)
	if err != nil {
		err = logger.Log.Errorf("err: %w", err)
		return err
	}

	// ステータスコードが200以外の場合、エラーにして返却
	if res.StatusCode != 200 {
		logger.Log.Debugf("batchIDが無効です/batchID: %s", batchID)
	}

	return nil
}