package orm

import "github.com/FourWD/middleware/orm"

type Maintenance struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleID  string `db:"vehicle_id" json:"vehicle_id" gorm:"type:varchar(45);"`
	LocationID string `db:"location_id" json:"location_id" gorm:"type:varchar(50);"`
	Mileage    string `db:"mileage" json:"mileage" gorm:"type:varchar(50);"`

	MaintenanceDate string `db:"maintenance_date" json:"maintenance_date" gorm:"type:varchar(100);"`
	Detail          string `db:"detail" json:"detail" gorm:"type:varchar(200);"`
}
