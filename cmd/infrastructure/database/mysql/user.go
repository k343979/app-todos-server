package mysql

import (
	"database/sql"

	"github.com/app-todos/cmd/domain/repository"
	"github.com/app-todos/cmd/domain/entity"
)

// ユーザ用DB構造体
type User struct {
	DB *sql.DB
}

// NewUser
// ユーザ用DB構造体の生成
// param db : DBコネクション
// return ユーザ用リポジトリインターフェース
func NewUser(db *sql.DB) repository.IUser {
	return &User{
		DB: db,
	}
}

func (u *User) ByID(id int) (*entity.User, error) {
	return &entity.User{}, nil
}

func (u *User) Update(user *entity.User) (*entity.User, error) {
	return user, nil
}