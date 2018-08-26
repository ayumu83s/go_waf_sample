package repository

// CRUDに関するI/Fを定義

type User interface {
	Get(id int)
}
