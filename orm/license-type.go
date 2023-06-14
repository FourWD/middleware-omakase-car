package orm

import "github.com/FourWD/middleware/orm"

type LicenseType struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `db:"name" json:"name" gorm:"type:varchar(45);"`
	NameEn string `db:"name_en" json:"name_en" gorm:"type:varchar(45);"`
}
