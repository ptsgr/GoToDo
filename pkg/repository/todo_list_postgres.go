package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/ptsgr/GoToDo"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list GoToDo.TodoList) (int, error) {
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

func (r *TodoListPostgres) GetAll(userID int) ([]GoToDo.TodoList, error) {
	var lists []GoToDo.TodoList
	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl
		INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1`,
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userID)
	return lists, err
}

func (r *TodoListPostgres) GetByID(userID, listID int) (GoToDo.TodoList, error) {
	var list GoToDo.TodoList
	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl
		INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userID, listID)
	return list, err
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s tl USING %s ul 
		WHERE tl.id = ul.list_id AND ul.user_id = $1 AND tl.id = $2`,
		todoListsTable, usersListsTable)

	_, err := r.db.Exec(query, userID, listID)
	return err
}

func (r *TodoListPostgres) Update(userID, listID int, input GoToDo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s tl SET %s FROM %s ul 
		WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d`,
		todoListsTable, setQuery, usersListsTable, argID, argID+1)

	args = append(args, listID, userID)

	_, err := r.db.Exec(query, args...)
	return err
}
