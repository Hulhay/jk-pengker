package repository

import (
	"gorm.io/gorm"
)

type repository struct {
	qry *gorm.DB
}
