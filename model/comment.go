package model

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Comment struct {
	ID     int64  `db:"comment_id" json:"id"`
	Body   string `json:"body"`
	TodoID int64  `json:"todo_id"`
}

func (c *Comment) Insert(tx *sqlx.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into comments (body, todo_id)
	values("hoge", 1)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.Body, c.TodoID)
}
