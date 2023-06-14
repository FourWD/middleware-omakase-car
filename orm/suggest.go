package orm

import "github.com/FourWD/middleware/orm"

type Suggest struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SuggestBy string `db:"suggest_by" json:"suggest_by" gorm:"index;default:null;type:varchar(100)"`
	Type      int    `db:"type" json:"type" gorm:"default:null;type:tinyint(2)"`
}
