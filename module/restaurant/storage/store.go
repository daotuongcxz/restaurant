package restaurantstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

type SqlStore struct {
	Db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func NewSQLStore1(db *gorm.DB) *SqlStore {
	return &SqlStore{Db: db}
}
