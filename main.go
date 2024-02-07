package main

import (
    // "database/sql"
    // "log"
    "main/server"
    // "github.com/gin-gonic/gin"

    "flag"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
	// コマンドライン引数の解析
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.Parse()

	// サーバーの起動
	app.Serve(addr)
}