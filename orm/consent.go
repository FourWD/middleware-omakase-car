package orm

import "github.com/FourWD/middleware/orm"

type Consent struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	RefCode     string `db:"ref_code" json:"ref_code" gorm:"type:varchar(50);"`
	Name        string `db:"name" json:"name" gorm:"type:varchar(100);"`
	Description string `db:"description" json:"description" gorm:"type:text;"`
	Status      uint64 `db:"status" json:"status" gorm:"type:int;"`
}
