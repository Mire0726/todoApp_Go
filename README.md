# ⚫︎ 概要

## 技術選定
<p>バックエンド：Golang</p>
<p>フロントエンド：next.js(localhost:3000)</p>
<p>DB：MySQL(phpmyadmin)(localhost:4000)</p>
<p>Docker</p>

サイバーエージェント社にてWomenGoCollegeでのインターンを終えたばかりだったので、Go言語をバックエンドの言語として利用し、DBはMySQLを利用しました。
フロントエンドに関しては、Golangと組み合わせて利用する記事をいくつか見たことがあったのでNext.jsを利用しました。
Dockerでは、バックエンドを除くDB(MySQL,phpmyadin),フロントエンドを構築しました

## API 仕様
###　フォルダ構成
backend>appフォルダの中でmodelフォルダとhandlerフォルダ,
serverフォルダに分けました。

modelフォルダ→クエリに直接指示を出す関数

handlerフォルダ→modelフォルダの関数を呼び出し、バックエンドのための処理を書く関数

serverフォルダには→httpサーバとの接続、ルーティング設定

### 各APIのルートと関数
タイトルを指定して TODO を作成するAPI

リクエストのルート：`e.POST("/todo", handler.HandleTodoCreate)`

todoのidはuuidを利用して、一意のidを割り当てている

作成した TODO の一覧を取得するAPI

リクエストのルート：`e.GET("/todos", handler.HandleGetAllTodos)`

指定した TODO を変更するAPI

リクエストのルート：`e.PUT("/todo", handler.HandleUpdateTodoByTitle)`

指定した TODO を削除するAPI

リクエストのルート：`e.DELETE("/todo/:id", handler.HandleDeleteTodo)`




## 起動、動作確認方法
### 起動
① ターミナルで`Docker compose up`を行い、Dockerを起動

②別ターミナルで`go run main.go`でバックエンドを起動

③curlコマンドでそれぞれのAPIの動作を確認

###　動作確認コマンド

タイトルを指定して TODO を作成するAPI：curlコマンドの例
```
curl -X POST http://localhost:8080/todo -H "Content-Type: application/json" -d '{"title": "New ToDo Item"}'
```

作成した TODO の一覧を取得するAPI
`curl -X GET http://localhost:8080/todos
`

指定した TODO を変更するAPI

```
curl -X PUT http://localhost:8080/todo -H "Content-Type: application/json" -d '{"currentTitle": "既存のタイトル", "newTitle": "新しいタイトル"}'

```
指定した TODO を削除するAPI

`curl -X DELETE http://localhost:8080/todo/"削除したいtodoのid"
`


# ⚫︎ 工夫したところ 
・一覧の取得とtodoの追加はwebのフロントからできるようにしました。（localhost:3000で、参照と追加が可能です。）

・POSTやPUTを使い分けました。

・modelとhandlerでフォルダを分ける構成にし、可読性をあげました。

# ⚫︎ 生成AI について
chatGPT 4.0を利用しました。

## 聞いたこと

### createdAtフィールドの生成方法
> go関数において、作成した日を自動で入れる機能を実装してください

timeパッケージを提示され、利用しました。

### エラーの解決方法

<p>timeパッケージに関するエラー</p>
>Error parsing createdAt: parsing time "2023-12-28 00:00:00" as "2006-01-02T15:04:05Z07:00": cannot parse " 00:00:00" as "T"

=>layoutを利用する方法を提示され、利用しました。以下参考にしたコードです（少し変更を加えました）

```
layout := "2006-01-02 15:04:05" // Goのレイアウト文字列をこの形式に合わせる
createdAt, err := time.Parse(layout, createdAtStr)
if err != nil {
    log.Printf("Error parsing createdAt: %v", err)
    // エラー処理
}

updatedAt, err := time.Parse(layout, updatedAtStr)
if err != nil {
log.Printf("Error parsing updatedAt: %v", err)
// エラー処理
}

```

### エラーメッセージの追加

>// HandleTodoCreate はToDo情報の作成処理を行います
>func HandleTodoCreate(c echo.Context) error {
>    req := &todoCreateRequest{}
>    if err := c.Bind(req); err != nil {
>        return err
>    }
>
>   // UUIDでTodoIDを生成
>    todoID, err := uuid.NewRandom()
>    if err != nil {
>        return err
>    }

>    // データベースにToDoデータを登録
>    if err := model.InsertTodo(&model.Todo{
>        ID:        todoID.String(),
>        Name:      req.Name,
>        AtCreate:  0,
>        AtUpdate:  0,
>    }); err != nil {
>        return err
>    }

>    log.Println("ToDo created with ID:", todoID.String())
>    // 生成したToDoIDをレスポンスとして返却
>    return c.JSON(http.StatusOK, echo.Map{"todoID": todoID.String()})
>
>細かいエラーメッセージを追加してください

=>log.Printfを利用したエラーメッセージを提示され、利用しました。
その後のコードもそれを参考にエラーメッセージを書きました。

### フロントエンドの実装
```
package handler

import (
	"log"
	"main/model"
	
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

// todoCreateRequest クライアントからのToDo作成リクエストの構造を定義
type todoCreateRequest struct {
    Title string `json:"title"`
}


// HandleTodoCreate クライアントがtodoを作成するための処理
func HandleTodoCreate(c echo.Context) error {
    req := &todoCreateRequest{}
    if err := c.Bind(req); err != nil {
        log.Printf("Error binding request: %v", err)
        return echo.NewHTTPError(http.StatusServiceUnavailable, "Error binding request")
    }

    todoID, err := uuid.NewRandom()
    if err != nil {
        log.Printf("Error generating UUID: %v", err)
        return echo.NewHTTPError(http.StatusServiceUnavailable, "Error generating UUID")
    }

    currentTime := time.Now()

    if err := model.InsertTodo(&model.Todo{
        ID:        todoID.String(),
        Title:     req.Title,
        CreatedAt: currentTime,
        UpdatedAt: currentTime,
    }); err != nil {
        log.Printf("Error inserting todo into database: %v", err)
        return echo.NewHTTPError(http.StatusServiceUnavailable, "Error inserting todo into database")
    }

    log.Println("ToDo created with ID:", todoID.String())
    return c.NoContent(http.StatusNoContent)
}


// HandleGetAllTodos 全てのToDoアイテムを取得
func HandleGetAllTodos(c echo.Context) error {
    todos, err := model.GetAllTodos()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve todos")
    }

    return c.JSON(http.StatusOK, todos)
}

type updateTodoRequest struct {
    CurrentTitle string `json:"currentTitle"`
    NewTitle     string `json:"newTitle"`
}

// HandleUpdateTodoByTitle　タイトルを指定して更新する処理
func HandleUpdateTodoByTitle(c echo.Context) error {
    req := new(updateTodoRequest)
    if err := c.Bind(req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
    }

    err := model.UpdateTodoByTitle(req.CurrentTitle, req.NewTitle)
    if err != nil {
        if err == model.ErrNoTodoFound {
            return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
        }
        // その他のエラーの場合は503を返す
        return echo.NewHTTPError(http.StatusServiceUnavailable, "Error updating todo")
    }

    return c.NoContent(http.StatusNoContent)
}

func HandleDeleteTodo(c echo.Context) error {
    id := c.Param("id")

    err := model.DeleteTodoByID(id)
    if err != nil {
        if err == model.ErrNoTodoFound {
            return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
        }
        log.Printf(err.Error())
        return echo.NewHTTPError(http.StatusServiceUnavailable, "Error deleting todo")
    }

    return c.NoContent(http.StatusNoContent)
}
```
>のバックエンドを前提にフロントのwebページを構築します。
>階層とコードの例を出力してください

# 参考にした記事
https://qiita.com/hrk_ym/items/c73c5ad41c92688c3b94

https://qiita.com/Senritsu420/items/413e0ffab3b92b60e5cb
<p>yamlファイルを参考、引用しました</p>

# ⚫︎ 反省と展望
フロントエンドに関しては、時間が足りず大部分を生成AIに頼ってしまいました。そのため、自分のコードになっていない感覚があります。

また、今回は負荷試験などを行うことができませんでした。vegeta,proffなどの負荷試験ツールはもともと気になっていたので、今後利用できたらなと思います。
