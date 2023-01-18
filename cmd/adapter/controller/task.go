package controller

import (
	"context"
	"strconv"

	"github.com/app-todos/cmd/infrastructure/database"
	"github.com/app-todos/cmd/infrastructure/database/mysql"
	"github.com/app-todos/cmd/usecase"
	"github.com/app-todos/library/logger"
	"github.com/gin-gonic/gin"
)

// タスク用コントローラ構造体
type Task struct {
	uc usecase.ITask
}

// タスク用コントローラインターフェース
type ITask interface {
	// Fetch
	// 該当条件を満たすタスク情報の取得
	// param *gin.Context
	Fetch(*gin.Context)

	// ByID
	// IDによるタスク情報の取得
	// param *gin.Context
	ByID(*gin.Context)

	// Update
	// タスク情報の更新
	// param *gin.Context
	Update(*gin.Context)
}

// NewTask
// タスク用コントローラ構造体の生成
// param ctx : コンテキスト
// return タスク用コントローラインターフェース
func NewTask(ctx context.Context) ITask {
	db, err := database.Connect()
	if err != nil {
		logger.Log.Errorf("DB Connect err: %w", err)
	}
	ITaskRepo := mysql.NewTask(ctx, db)
	return &Task{
		uc: usecase.NewTask(ITaskRepo),
	}
}

func (t *Task) Fetch(c *gin.Context) {
	// 取得処理の実行
	tasks, err := t.uc.Fetch()
	if err != nil {
		logger.Log.Errorf("func Fetch err: %w", err)
	}
	// レスポンス設定
	c.JSON(200, gin.H{"tasks": tasks})
}

func (t *Task) ByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log.Errorf("Convert ID err: %w", err)
	}
	// 取得処理の実行
	task, err := t.uc.ByID(id)
	if err != nil {
		logger.Log.Errorf("func ByID err: %w", err)
	}
	// レスポンス設定
	c.JSON(200, gin.H{"task": task})
}

func (t *Task) Update(c *gin.Context)
