package skey

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	db *gorm.DB
	// trans  ITransService
	// LangID lang.LangID
}

func New(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}
