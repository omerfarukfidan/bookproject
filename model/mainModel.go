package model

type MainModel struct {
	user UserModel `json:"user"`
	book BookModel `json:"book"`
}
