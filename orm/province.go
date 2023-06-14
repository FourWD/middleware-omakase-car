package orm

import "github.com/FourWD/middleware/orm"

type Province struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	Code string `db:"code" json:"code" gorm:"not null;index;unique;type:varchar(2)"`

	Name     string `db:"name" json:"name" gorm:"not null;index;type:varchar(150)"`
	NameEn   string `db:"name_en" json:"name_en" gorm:"not null;index;type:varchar(150)"`
	RegionID int    `db:"region_id" json:"region_id" gorm:"default:0;not null;index;type:int(5)"`
}
