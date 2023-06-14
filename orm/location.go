package orm

import "github.com/FourWD/middleware/orm"

type Location struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	LocationGroupID string `db:"location_group_id" json:"location_group_id" gorm:"type:varchar(150);"`

	Name string `db:"name" json:"name" gorm:"type:varchar(150);"`
}
