// タスク用リポジトリパッケージ
package repository

import (
	"github.com/app-todos/cmd/domain/entity"
)

// タスク用リポジトリインターフェース
type ITask interface {
	// Fetch
	// 該当条件を満たすタスク情報の取得
	// return タスク情報のスライス
	// return エラー情報
	Fetch() ([]*entity.Task, error)

	// ByID
	// IDによるタスク情報の取得
	// param id : タスクID
	// return タスク情報
	// return エラー情報
	ByID(id int) (*entity.Task, error)

	// Update
	// タスク情報の更新
	// param task : タスク更新内容
	// return 更新後タスク情報
	// return エラー情報
	Update(task *entity.Task) (*entity.Task, error)
}
