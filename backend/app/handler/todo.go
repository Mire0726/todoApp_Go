
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
