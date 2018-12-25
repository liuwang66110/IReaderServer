package models

import "time"

type User struct {
	ID          uint64    `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string    `gorm:"type:varchar(64);not null;default:'';unique_index:uk_name"`
	Mobile      uint64    `gorm:"not null; default:0;unique_index:uk_mobile" json:",omitempty"`
	Password    string    `gorm:"type:varchar(60); not null; default:''" json:"-"`
	Status      int       `gorm:"not null; default : 1"`
	Token       string    `gorm:"type:varchar(40);not null; default:'' ; unique_index:token" json:",omitempty"`
	ExpiredAt   time.Time `gorm:"not null; default: '2000-01-01 00:00:00'" json:"-"`
	CreatedAt   time.Time `gorm:"not null; default: '2000-01-01 00:00:00'" json:"-"`
	UpdatedAt   time.Time `gorm:"not null;" sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"  json:"-"`
}

const (
	USER_ST_ENABLE  = 1
	USER_ST_DISABLE = 2
	USER_ST_DELETED = 3
)

//const (
//    USER_LOGIN_TOKEN = 1
//    USER_LOGIN_PASSWORD = 2
//)

const USER_TOKEN_DURATION = time.Hour * 24 * 300

func (User) TableName() string {
	return "user_main"
}
