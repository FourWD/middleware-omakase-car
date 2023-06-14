package orm

import (
	"github.com/FourWD/middleware/orm"
)

type Grade struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(25); unique_index " `
	Img      string `db:"img" json:"img" gorm:"type:varchar(255);"`
	IsActive bool   `db:"is_active" json:"is_active" gorm:"type:bool;"`
}
