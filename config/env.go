// 環境変数設定ファイル
package config

import (
	"os"
	"sync"
)

// 初回読み込み時のみ取得(レースコンディション対策)
var once sync.Once

// DB情報構造体
var DB struct {
	Driver   string
	Host     string
	Name     string
	User     string
	Password string
	Port     string
	Timezone string
}

// 初期処理
func init() {
	loadEnv()
}

// loadEnv
// envファイルを読み込み、構造体に環境変数を設定
func loadEnv() {
	once.Do(newDBConfig)
}

// newDBConfig
// 環境ファイルから取得したDB情報を構造体に設定
func newDBConfig() {
	DB.Driver = os.Getenv("DB_DRIVER")
	DB.Host = os.Getenv("DB_HOST")
	DB.Name = os.Getenv("DB_DATABASE")
	DB.User = os.Getenv("DB_USERNAME")
	DB.Password = os.Getenv("DB_PASSWORD")
	DB.Port = os.Getenv("DB_PORT")
	DB.Timezone = os.Getenv("TIMEZONE")
}
