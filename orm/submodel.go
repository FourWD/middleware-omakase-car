package orm

import "github.com/FourWD/middleware/orm"

type SubModel struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name    string `db:"name" json:"name" gorm:"index;not null;type:varchar(50)"`
	BrandID int    `db:"brand_id" json:"brand_id" gorm:"not null;type:int(11)"`
	ModelID int    `db:"model_id" json:"model_id" gorm:"not null;type:int(11)"`
	Key     string `db:"key" json:"key" gorm:"index;not null;type:varchar(45)"`
}
