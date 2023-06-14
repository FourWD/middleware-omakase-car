package orm

import (
	"github.com/FourWD/middleware/orm"
)

type Vehicle struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	License string `json:"license" db:"license"`
}
