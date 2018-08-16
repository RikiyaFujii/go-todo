package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/voyagegroup/go-todo/model"

	"github.com/jmoiron/sqlx"
)

// Todo はTodoへのリクエストに関する制御をします
type Todo struct {
	DB *sqlx.DB
}

func (t *Todo) Order(w http.ResponseWriter, r *http.Request) error {
	order := r.FormValue("order")
	todos, err := model.Order(t.DB, order)
	if err != nil {
		return err
	}
	return JSON(w, 200, todos)
}

// 検索
func (t *Todo) Search(w http.ResponseWriter, r *http.Request) error {
	title := r.FormValue("title")
	completed := r.FormValue("completed")
	// createdAt := r.FormValue("created_at")
	c, err := strconv.ParseBool(completed)
	todos, err := model.Search(t.DB, title, c)
	if err != nil {
		return err
	}
	return JSON(w, 200, todos)
}

// GetはDBからユーザを取得して結果を返します
func (t *Todo) Get(w http.ResponseWriter, r *http.Request) error {
	todos, err := model.TodosAll(t.DB)
	if err != nil {
		return err
	}
	return JSON(w, 200, todos)
}

func (t *Todo) Put(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		result, err := todo.Update(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		todo.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusOK, todo)
}

// PostはタスクをDBに追加します
// todoをJSONとして受け取ることを想定しています。
func (t *Todo) Post(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		result, err := todo.Insert(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		todo.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusCreated, todo)
}

func (t *Todo) Delete(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		_, err := todo.Delete(tx)
		if err != nil {
			return err
		}
		return tx.Commit()
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusOK, todo)
}

func (t *Todo) Toggle(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		result, err := todo.Toggle(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		todo.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusOK, nil)
}
