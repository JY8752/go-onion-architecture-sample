package infrastructure

import (
	"log"
	"time"

	"github.com/JY8752/go-onion-architecture-sample/domain/model"
	"github.com/JY8752/go-onion-architecture-sample/domain/repository"
	db "github.com/JY8752/go-onion-architecture-sample/infrastructure"
)

type todoRepository struct {
	dbClient *db.DBClient
}

func NewTodoRepository(db *db.DBClient) repository.TodoRepository {
	return &todoRepository{
		dbClient: db,
	}
}

func (t *todoRepository) Create(userId model.UserId, title model.Title, description model.Description) (model.TodoId, error) {
	stmt, err := t.dbClient.Client.Prepare("INSERT INTO todos (user_id, title, description, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(userId, title, description, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return model.TodoId(id), nil
}

func (t *todoRepository) List(id model.UserId) ([]model.Todo, error) {
	stmt, err := t.dbClient.Client.Prepare("SELECT id, title, description, created_at FROM todos WHERE user_id = ? AND delete_at IS NULL")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt)
		if err != nil {
			log.Printf("err: %s\n", err.Error())
			continue
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *todoRepository) Delete(id model.TodoId) error {
	stmt, err := t.dbClient.Client.Prepare("UPDATE todos SET delete_at = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
