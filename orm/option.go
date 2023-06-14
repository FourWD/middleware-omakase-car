package orm

import "github.com/FourWD/middleware/orm"

type Option struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Detail string  `db:"detail" json:"detail"  gorm:"default:null;type:varchar(45)"`
	Price  float32 `db:"price" json:"price"  gorm:"default:null;type:float(11,2)"`
}
