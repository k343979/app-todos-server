package logger

import (
	"io/ioutil"
	"strings"

	log "github.com/cihub/seelog"
)

// Logger
var Log log.LoggerInterface

// Set
// ログの初期設定
// param processName : 処理の種別
func Set(processName string) {
	// 処理種別が空の場合、panic
	if processName == "" {
		panic("kind of process is unclear")
	}

	// log.xmlファイルの読み取り
	buf, err := ioutil.ReadFile("/go/src/github.com/app-todos/log.xml")
	if err != nil {
		panic(err)
	}

	// $PROCESS$を処理種別に書き換え
	xmldoc := strings.Replace(string(buf), "$PROCESS$", processName, 1)
	// log.xmlの設定をもとにloggerを作成
	logger, err := log.LoggerFromConfigAsBytes([]byte(xmldoc))
	if err != nil {
		panic(err)
	}

	Log = logger
}
