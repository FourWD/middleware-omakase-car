package orm

import "github.com/FourWD/middleware/orm"

type Menu struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Path    string `db:"path" json:"path" gorm:"type:varchar(255);"`
	Name    string `db:"name" json:"name" gorm:"type:varchar(45);"`
	GroupID uint   `db:"group_id" json:"group_id" gorm:"type:int;"`
}
