package controller

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/go-todo/model"
	"net/http"
)

type Comment struct {
	DB *sqlx.DB
}

func (c *Comment) Post(w http.ResponseWriter, r *http.Request) error {
	var comment model.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return err
	}

	if err := TXHandler(c.DB, func(tx *sqlx.Tx) error {
		result, err := comment.Insert(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		comment.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusCreated, comment)
}
