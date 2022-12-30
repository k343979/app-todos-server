// ユーザ用エンティティパッケージ
package entity

// ユーザ構造体
// usersテーブルのレコード
type User struct {
	ID        int    `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	Tel       string `json:"phoneNumber"`
	JobTitle  string `json:"jobTitle"`
}