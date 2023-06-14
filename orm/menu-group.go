package orm

import "github.com/FourWD/middleware/orm"

type MenuGroup struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(45);"`
}
