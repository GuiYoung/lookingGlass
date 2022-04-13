package accessibleLG

import "gorm.io/gorm"

type accessibleLG struct {
	gorm.Model
	LgUrl    string `gorm:"lg_url"`
	LgIsp    string `json:"isp" gorm:"lg_isp"`
	LgAS     int    `json:"as" gorm:"lg_AS"`
	LgStatus int    `gorm:"lg_status"`
}

func InsertLgUrl(lg *accessibleLG) (err error) {
	if err = Db.Create(lg).Error; err != nil {
		return
	}
	return
}
