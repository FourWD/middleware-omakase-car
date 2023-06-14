package orm

import (
	"github.com/FourWD/middleware/orm"
)

type User struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Fullname string `db:"fullname" json:"fullname"`
}
