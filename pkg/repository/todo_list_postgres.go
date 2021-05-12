package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ptsgr/GoToDo"
)

type TodoLIstPostgres struct {
	db *sqlx.DB
}

func NewTodoLIstPostgres(db *sqlx.DB) *TodoLIstPostgres {
	return &TodoLIstPostgres{db: db}
}

func (r *TodoLIstPostgres) Create(userID int, list GoToDo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
