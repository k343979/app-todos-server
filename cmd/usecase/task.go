// タスク用ユースケースパッケージ
package usecase

import (
	"github.com/app-todos/cmd/domain/entity"
	"github.com/app-todos/cmd/domain/repository"
)

// タスク用ユースケース構造体
type Task struct {
	Repo repository.ITask
}

// タスク用ユースケースインターフェース
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
	Update(req map[string]any) (*entity.Task, error)
}

// NewTask
// タスク用ユースケース構造体の生成
// param r : タスク用リポジトリインターフェース
// return タスク用ユースケースインターフェース
func NewTask(r repository.ITask) ITask {
	return &Task{
		Repo: r,
	}
}

func (t *Task) Fetch() ([]*entity.Task, error) {
	return t.Repo.Fetch()
}

func (t *Task) ByID(id int) (*entity.Task, error) {
	return t.Repo.ByID(id)
}

func (t *Task) Update(req map[string]any) (*entity.Task, error) {
	task := &entity.Task{
		ID:    int(req["id"].(float64)),
		Title: req["title"].(string),
		Label: req["label"].(string),
	}
	return t.Repo.Update(task)
}
