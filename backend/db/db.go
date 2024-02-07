package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


const driverName = "mysql"

var Conn *sql.DB

func init() {
	// 接続情報
	user := "root"                      // MySQLのユーザー名
	password := "go-college"            // MySQLのパスワード
	host := "127.0.0.1"                 // MySQLがリッスンしているホスト (ローカルホスト)
	port := "3307"                      // MySQLがリッスンしているポート
	database := "todo_app"              // 使用するデータベース名

	var err error
	Conn, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
	if err := Conn.Ping(); err != nil {
		log.Fatalf("can't connect to mysql server. "+
			"MYSQL_USER=%s, "+
			"MYSQL_PASSWORD=%s, "+
			"MYSQL_HOST=%s, "+
			"MYSQL_PORT=%s, "+
			"MYSQL_DATABASE=%s, "+
			"error=%+v",
			user, password, host, port, database, err)
	}
}
