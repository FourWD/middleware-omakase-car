package orm

import (
	"github.com/FourWD/middleware/orm"
)

type Finance struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name             string `db:"name" json:"name" gorm:"type:varchar(50); unique_index " `
	ContactName      string `db:"contact_name" json:"contact_name" gorm:"type:varchar(50);"`
	ContactTel       string `db:"contact_tel" json:"contact_tel" gorm:"type:varchar(50);"`
	ImgSize1         string `db:"img_size1" json:"img_size1" gorm:"type:varchar(100);"`
	ImgSize2         string `db:"img_size2" json:"img_size2" gorm:"type:varchar(100);"`
	NameTh           string `db:"name_th" json:"name_th" gorm:"type:varchar(150);"`
	OrderBy          uint   `db:"order_by" json:"order_by" gorm:"type:int;"`
	IsActive         uint   `db:"is_active" json:"is_active" gorm:"type:int;"`
	PrefinanceStatus uint   `db:"prefinance_status" json:"prefinance_status" gorm:"type:int;"`
}
