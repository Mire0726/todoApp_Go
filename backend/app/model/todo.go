package model

import (
	// "database/sql"
	"db"
	"log"
	"time"
	"errors"
)

// todoテーブルデータ
type Todo struct {
	ID        string
	Title      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// InsertTodo データベースをレコードを登録する
func InsertTodo(record *Todo) error {
	if _, err := db.Conn.Exec(
		"INSERT INTO todo_app (id, title, createdAt,updatedAt ) values (?, ?, ?, ?)",
		record.ID,
		record.Title,
		record.CreatedAt,
		record.UpdatedAt,
	); err != nil {
		log.Printf("Error inserting todo into database(at model pkg): %v", err)
		return err
	}
	return nil
}

// GetAllTodos データベースから全てのToDoを取得する
func GetAllTodos() ([]Todo, error) {
	rows, err := db.Conn.Query("SELECT id, title, createdAt, updatedAt FROM todo_app")
	if err != nil {
		log.Printf("Error querying todos from database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		var createdAtStr, updatedAtStr string // 中間変数を用意

		if err := rows.Scan(&todo.ID, &todo.Title, &createdAtStr, &updatedAtStr); err != nil {
			log.Printf("Error scanning todo: %v", err)
			return nil, err
		}

		layout := "2006-01-02 15:04:05"
		// createdAt と updatedAt を time.Time 型に変換
		todo.CreatedAt, err = time.Parse(layout, createdAtStr)
		if err != nil {
			log.Printf("Error parsing createdAt: %v", err)
			return nil, err
		}

		todo.UpdatedAt, err = time.Parse(layout, updatedAtStr)
		if err != nil {
			log.Printf("Error parsing updatedAt: %v", err)
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}

	return todos, nil
}

//５０３を返す
var ErrNoTodoFound = errors.New("no todo found")

// UpdateTodoByTitle
func UpdateTodoByTitle(currentTitle, newTitle string) error {
	updatedTime :=time.Now()

    // 最初にマッチしたレコードを更新
    result, err := db.Conn.Exec("UPDATE todo_app SET title = ?, updatedAt=? WHERE title = ? LIMIT 1", 
	newTitle,
	updatedTime,
	currentTitle,
)
    if err != nil {
        log.Printf("Error updating todo: %v", err)
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Printf("Error getting rows affected: %v", err)
        return err
    }

    if rowsAffected == 0 {
        return ErrNoTodoFound
    }

    return nil
}


func DeleteTodoByID(id string) error {
    result, err := db.Conn.Exec("DELETE FROM todo_app WHERE id = ?", id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return ErrNoTodoFound 
    }

    return nil
}

