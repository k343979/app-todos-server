// DBの接続設定用パッケージ
package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/app-todos/config"
	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

// Connect
// DB接続
// return DBコネクション
// return エラー情報
func Connect() (*bun.DB, error) {
	// タイムゾーン設定
	jst, err := time.LoadLocation(config.DB.Timezone)
	if err != nil {
		return nil, err
	}

	// DB接続情報を設定
	c := &mysql.Config{
		DBName:    config.DB.Name,
		User:      config.DB.User,
		Passwd:    config.DB.Password,
		Addr:      fmt.Sprintf("%s:%s", config.DB.Driver, config.DB.Port),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	// DB接続開始
	sqldb, err := sql.Open(config.DB.Driver, c.FormatDSN())
	if err != nil {
		return nil, err
	}

	// bunにDBコネクションを設定
	db := bun.NewDB(sqldb, mysqldialect.New())
	return db, nil
}
