package orm

import "github.com/FourWD/middleware/orm"

type Amphor struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Code       string `db:"code" json:"code" gorm:"type:varchar(4); unique_index" `
	Name       string `db:"name" json:"name" gorm:"type:varchar(150);"`
	NameEn     string `db:"name_en" json:"name_en" gorm:"type:varchar(150);"`
	ProvinceID uint   `db:"provinceid" json:"provinceid" gorm:"type:int;"`
}
