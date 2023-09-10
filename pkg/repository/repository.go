package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/poroshin-ao/test-project"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	UpdateList(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, itemId int) (todo.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
		TodoList:      newTodoListPostgres(db),
		TodoItem:      newTodoItemPostgres(db),
	}
}
