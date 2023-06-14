package orm

import "github.com/FourWD/middleware/orm"

type EngineOil struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name    string `db:"name" json:"name"`
	NameEng string `db:"name_eng" json:"name_eng"`
}