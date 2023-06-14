package orm

import "github.com/FourWD/middleware/orm"

type Model struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"not null;type:varchar(50)"`
	BrandID  int    `db:"brand_id" json:"brand_id" gorm:"not null;type:int(11)"`
	ModelKey string `db:"model_key" json:"model_key"  gorm:"not null;type:varchar(45)"`
}
