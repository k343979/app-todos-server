// タスク用DB処理パッケージ
package mysql

import (
	"context"

	"github.com/app-todos/cmd/domain/entity"
	"github.com/app-todos/cmd/domain/repository"
	"github.com/uptrace/bun"
)

// タスク用DB構造体
type Task struct {
	DB  *bun.DB
	ctx context.Context
}

// NewTask
// タスク用DB構造体の生成
// param ctx : コンテキスト
// param db : DBコネクション
// return タスク用リポジトリインターフェース
func NewTask(ctx context.Context, db *bun.DB) repository.ITask {
	return &Task{
		DB:  db,
		ctx: ctx,
	}
}

func (t *Task) Fetch() ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := t.DB.NewSelect().
		Model(tasks).
		Scan(t.ctx); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *Task) ByID(id int) (*entity.Task, error) {
	var task = &entity.Task{}
	if err := t.DB.NewSelect().
		Model(task).
		Where("id = ?", id).
		Scan(t.ctx); err != nil {
		return nil, err
	}
	return task, nil
}

func (t *Task) Update(task *entity.Task) (*entity.Task, error) {
	res, err := t.DB.NewUpdate().
		Model(task).
		Where("id = ?", task.ID).
		Exec(t.ctx)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return nil, nil
	}
	return t.ByID(task.ID)
}
