package accessibleLG

import "gorm.io/gorm"

type lg struct {
	gorm.Model
	LgUrl string `gorm:"type:varchar(200);not null" json:"lgUrl"`
}

func InsertLgUrl(lg *lg) (err error) {
	if err = Db.Create(lg).Error; err != nil {
		return
	}
	return
}
