package orm

import "github.com/FourWD/middleware/orm"

type District struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ZipCode  uint   `db:"zip_code" json:"zip_code" gorm:"type:int; " `
	Name     string `db:"name" json:"name" gorm:"type:varchar(150);"`
	NameEn   string `db:"name_en" json:"name_en" gorm:"type:varchar(150);"`
	AmphorID uint   `db:"amphorid" json:"amphorid" gorm:"type:int;"`
}
