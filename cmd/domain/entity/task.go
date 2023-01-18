// タスク用エンティティパッケージ
package entity

// タスク構造体
// tasksテーブルのレコード
type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Label   string `json:"label"`
	Done_at string `json:"doneAt"`
}
