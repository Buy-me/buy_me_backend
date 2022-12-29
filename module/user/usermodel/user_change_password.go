package usermodel

type UserChangePassword struct {
	NewPassword string `json:"new_password" gorm:"-"`
	OldPassword string `json:"old_password" gorm:"-"`
}

func (UserChangePassword) TableName() string {
	return User{}.TableName()
}
