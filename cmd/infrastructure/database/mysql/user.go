// ユーザ用DB処理パッケージ
package mysql

import (
	"context"

	"github.com/app-todos/cmd/domain/entity"
	"github.com/app-todos/cmd/domain/repository"
	"github.com/uptrace/bun"
)

// ユーザ用DB構造体
type User struct {
	DB  *bun.DB
	ctx context.Context
}

// NewUser
// ユーザ用DB構造体の生成
// param ctx : コンテキスト
// param db : DBコネクション
// return ユーザ用リポジトリインターフェース
func NewUser(ctx context.Context, db *bun.DB) repository.IUser {
	return &User{
		DB:  db,
		ctx: ctx,
	}
}

func (u *User) ByID(id int) (*entity.User, error) {
	var user = &entity.User{}
	if err := u.DB.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(u.ctx); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Update(user *entity.User) (*entity.User, error) {
	res, err := u.DB.NewUpdate().
		Model(user).
		Where("id = ?", user.ID).
		Exec(u.ctx)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return nil, nil
	}
	return u.ByID(user.ID)
}
