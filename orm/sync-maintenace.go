package orm

import "github.com/FourWD/middleware/orm"

type SyncMaintenace struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Status int    `db:"status" json:"status" gorm:"default:null;type:int(1)"`
	Source string `db:"source" json:"source" gorm:"index;default:null;type:varchar(6)"`
	ApiUrl string `db:"api_url" json:"api_url" gorm:"type:varchar(255)"`
}
