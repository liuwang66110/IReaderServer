package models

import "time"

type UserInfo struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint64    `gorm:"not null;unique_index:uk_usd"`
	Name      string    `gorm:"type:varchar(64);not null;default:'';"`
	Mobile    string    `gorm:"not null; default:'';" json:",omitempty"`
	Gender    int       `gorm:"not null; default:3;"`
	Email     string    `gorm:"type:varchar(64);not null;default:'';"`
	UpdatedAt time.Time `gorm:"not null;" sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"  json:"-"`
}

const (
	USER_GENDER_MALE = 1 // 男性
	USER_GENDER_FEMALE = 2 // 女性
	USER_GENDER_NONE = 3 // 女性
)

func (UserInfo) TableName() string {
	return "user_info"
}
