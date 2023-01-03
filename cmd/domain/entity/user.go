// ユーザ用エンティティパッケージ
package entity

// ユーザ構造体
// usersテーブルのレコード
type User struct {
	ID        int    `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	JobTitle  string `json:"jobTitle"`
}
