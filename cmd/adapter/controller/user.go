// ユーザ用コントローラーパッケージ
package controller

import (
	"strconv"
	"database/sql"

	"github.com/app-todos/library/logger"
	"github.com/app-todos/cmd/usecase"
	"github.com/app-todos/cmd/infrastructure/database/mysql"

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
// return ユーザ用コントローラインターフェース
func NewUser() IUser {
	IUserRepo := mysql.NewUser(&sql.DB{})
	return &User{
		uc: usecase.NewUser(IUserRepo),
	}
}

func (u *User) ByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := u.uc.ByID(id)
	c.JSON(200, gin.H{"user": user})
}

func (u *User) Update(c *gin.Context) {
	// リクエストパラメータの取得
	req := make(map[string]any)
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error(err)
	}
	user, _ := u.uc.Update(req)
	c.JSON(200, gin.H{"user": user})
}