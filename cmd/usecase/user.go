// ユーザ用ユースケースパッケージ
package usecase

import (
	"github.com/app-todos/cmd/domain/entity"
	"github.com/app-todos/cmd/domain/repository"
)

// ユーザ用ユースケース構造体
type User struct {
	Repo repository.IUser
}

// ユーザ用ユースケースインターフェース
type IUser interface {
	// ByID
	// IDによるユーザ情報の取得
	// param id : ユーザID
	// return ユーザ情報
	// return エラー情報
	ByID(id int) (*entity.User, error)

	// Update
	// ユーザ情報の更新
	// param req : ユーザ更新内容
	// return 更新後ユーザ情報
	// return エラー情報
	Update(req map[string]any) (*entity.User, error)
}

// NewUser
// ユーザ用ユースケース構造体の生成
// param r : ユーザ用リポジトリインターフェース
// return ユーザ用ユースケースインターフェース
func NewUser(r repository.IUser) IUser {
	return &User{
		Repo: r,
	}
}

func (u *User) ByID(id int) (*entity.User, error) {
	return u.Repo.ByID(id)
}

func (u *User) Update(req map[string]any) (*entity.User, error) {
	user := &entity.User{
		ID:        int(req["id"].(float64)),
		LastName:  req["lastName"].(string),
		FirstName: req["firstName"].(string),
		Email:     req["email"].(string),
		Tel:       req["tel"].(string),
		JobTitle:  req["jobTitle"].(string),
	}
	return u.Repo.Update(user)
}
