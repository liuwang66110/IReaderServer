package models

type UserFriend struct {
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint64 `gorm:"not null; unique_index:uk_fid"`
	FriendID  uint64 `gorm:"not null; unique_index:uk_fid"`
	Status    uint64 `gorm:"not null;" json:"-"`
	Content   string `gorm:"type:varchar(64);not null;default:'';"`
	ConfirmAt StdTime `gorm:"not null;"`
	UpdatedAt StdTime `json:"-" sql:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

const (
	FRIEND_ST_WATTING = 1; // 向我申请的
	FRIEND_ST_REFUSE = 2; // 我拒绝的
	FRIEND_ST_AGREE = 3; // 已同意(好友)
	FRIEND_ST_DELETE = 4; // 已删除
)

func (UserFriend) TableName() string {
	return "user_friend"
}
