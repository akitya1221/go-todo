package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"main/config"

	"github.com/google/uuid"

	// SQLite3のドライバをインポート
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

// 作成するテーブル名を宣言
const (
	tableNameUser = "users"
)

// テーブルを作成する関数
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		SQLコマンド
		usersテーブルが存在しない場合、usersテーブルを作成する
	*/
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	// cmdUを起動
	Db.Exec(cmdU)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
