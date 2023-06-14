package orm

import "github.com/FourWD/middleware/orm"

type PermissionMenu struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PermissionID int `db:"permission_id" json:"permission_id" gorm:"default:null;type:int(11)" `
	MenuID       int `db:"menu_id" json:"menu_id" gorm:"default:null;type:int(11)"`
	ActionType   int `db:"action_type" json:"action_type"  gorm:"default:null;type:tinyint(1)"`
}
