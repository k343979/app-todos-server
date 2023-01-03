// ユーザ用リポジトリパッケージ
package repository

import (
	"github.com/app-todos/cmd/domain/entity"
)

// ユーザ用リポジトリインターフェース
type IUser interface {
	// ByID
	// IDによるユーザ情報の取得
	// param id : ユーザID
	// return ユーザ情報
	// return エラー情報
	ByID(id int) (*entity.User, error)

	// Update
	// ユーザ情報の更新
	// param user : ユーザ更新内容
	// return 更新後ユーザ情報
	// return エラー情報
	Update(user *entity.User) (*entity.User, error)
}