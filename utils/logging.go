package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(LogFile string) {
	// ファイルの読み込み | 作成 | 追記の設定
	logfile, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// フォーマットの設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// ログの出力先設定
	log.SetOutput(multiLogFile)
}
