package controller

import (
	"net/http"
)

type User struct {
	//DB *sqlx.DB
}

type UserModel struct {
	Name string `json:"name"`
}

// GetはDBからユーザを取得して結果を返します
func (t *User) Get(w http.ResponseWriter, r *http.Request) error {
	users := []*UserModel{
		&UserModel{Name: "hoge"},
	}
	return JSON(w, 200, users)
}
