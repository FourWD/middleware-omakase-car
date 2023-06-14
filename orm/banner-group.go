package orm

import "github.com/FourWD/middleware/orm"

type BannerGroup struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Code string `db:"code" json:"code" gorm:"unique;index;default:null;type:varchar(45)"`
	Name string `db:"name" json:"name" gorm:"not null;type:varchar(50)"`
}
