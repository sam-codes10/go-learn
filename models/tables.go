package models

type UserProfile struct {
	UserId      string `gorm:"primaryKey; column:user_id"`
	UserName    string `gorm:"column:user_name"`
	Name        string `gorm:"column:name"`
	Email       string `gorm:"column:email"`
	PhoneNumber string `gorm:"column:phone_number"`
	Password    string `gorm:"column:password"`
	// Verified    bool   `gorm:"column:verified"`
	Role      string `gorm:"column:role"`
	Lock      bool   `gorm:"column:lock"`
	CreatedAt int64  `gorm:"column:created_at"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
