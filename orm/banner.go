package orm

import "github.com/FourWD/middleware/orm"

type Banner struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BannerGroupID int    `json:"banner_grounp_id" query:"banner_grounp_id" db:"banner_grounp_id" gorm:"default:null;type:int(11)"`
	ImagePath     string `json:"image_path" query:"image_path" db:"image_path" gorm:"default:null;type:varchar(100)"`
	Url           string `json:"url" query:"url" db:"url" gorm:"default:null;type:varchar(100)"`
	YoutubeIframe string `json:"youtube_iframe" query:"youtube_iframe" db:"youtube_iframe" gorm:"default:null;type:varchar(100)"`
	Description   string `json:"description" query:"description" db:"description" gorm:"default:null;type:varchar(100)"`
}
