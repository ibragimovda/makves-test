package pgrepo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"makves/internal/entity"
)

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) GetByIds(ids []int) (items []entity.Item, err error) {
	query, args, err := sqlx.In(`select * from ueba where id in (?)`, ids)
	query = sqlx.Rebind(2, query)
	fmt.Println(query, args, err)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&items, query, args...)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
