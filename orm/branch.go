package orm

import "github.com/FourWD/middleware/orm"

type Branch struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name           string `db:"name" json:"name" gorm:"type:varchar(45);"`
	BranchInit     string `db:"branch_init" json:"branch_init" gorm:"type:varchar(45);"`
	HighlightColor string `db:"highlight_color" json:"highlight_color" gorm:"type:varchar(10);"`
}
