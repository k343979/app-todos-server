// ユーザ用コントローラーパッケージ
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

// ユーザ用コントローラ構造体
type User struct {
	uc usecase.IUser
}

// ユーザ用コントローラインターフェース
type IUser interface {
	// ByID
	// IDによるユーザ情報の取得
	// param *gin.Context
	ByID(*gin.Context)

	// Update
	// ユーザ情報の更新
	// param *gin.Context
	Update(*gin.Context)
}

// NewUser
// ユーザ用コントローラ構造体の生成
// param ctx : コンテキスト
// return ユーザ用コントローラインターフェース
func NewUser(ctx context.Context) IUser {
	db, err := database.Connect()
	if err != nil {
		logger.Log.Errorf("DB Connect err: %w", err)
	}
	IUserRepo := mysql.NewUser(ctx, db)
	return &User{
		uc: usecase.NewUser(IUserRepo),
	}
}

func (u *User) ByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log.Error(err)
	}
	// 取得処理の実行
	user, err := u.uc.ByID(id)
	if err != nil {
		logger.Log.Error(err)
	}
	// レスポンス設定
	c.JSON(200, gin.H{"user": user})
}

func (u *User) Update(c *gin.Context) {
	// リクエストパラメータの取得
	req := make(map[string]any)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error(err)
	}
	// 更新処理の実行
	user, err := u.uc.Update(req)
	if err != nil {
		logger.Log.Error(err)
	}
	// レスポンス設定
	c.JSON(200, gin.H{"user": user})
}
