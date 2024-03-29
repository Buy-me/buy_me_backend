package usermodel

import (
	"food_delivery/common"
	"food_delivery/module/address/addressmodel"
	"food_delivery/module/card/cardmodel"
	"food_delivery/module/food/foodmodel"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string                 `json:"email" gorm:"column:email;"`
	Password        string                 `json:"-" gorm:"column:password;"`
	Salt            string                 `json:"-" gorm:"column:salt;"`
	LastName        string                 `json:"last_name" gorm:"column:last_name;"`
	FirstName       string                 `json:"first_name" gorm:"column:first_name;"`
	Gender          string                 `json:"gender" gorm:"column:gender;"`
	BirthDate       string                 `json:"birth_date" gorm:"column:birth_date;"`
	Phone           string                 `json:"phone" gorm:"column:phone;"`
	Role            string                 `json:"role" gorm:"column:role;type:ENUM('user', 'admin')"`
	Avatar          *common.Image          `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	ListFoodInCart  []foodmodel.Food       `json:"items" gorm:"many2many:carts;"`
	ListFavourite   []foodmodel.Food       `json:"favourites" gorm:"many2many:favourites;"`
	ListAddress     []addressmodel.Address `json:"addresses"`
	ListCard        []cardmodel.Card       `json:"cards"`
}

func (u *User) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DbTypeUser)
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

// type Account struct {
// 	AccessToken  *tokenprovider.Token `json:"access_token"`
// 	RefreshToken *tokenprovider.Token `json:"refresh_token"`
// }

// func NewAccount(at, rt *tokenprovider.Token) *Account {
// 	return &Account{
// 		AccessToken:  at,
// 		RefreshToken: rt,
// 	}
// }

//func (u *User) ComparePassword(hasher common.Hasher) bool {
//	hashedPassword := hasher.Hash()
//	return u.Password == hashed
//}

func (u *User) IsActive() bool {
	if u == nil {
		return false
	}
	return u.Status == 1
}

//func (u *User) ToSimpleUser() *common.SimpleUser {
//	var simpleUser common.SimpleUser
//	simpleUser.ID = u.ID
//	simpleUser.Email = u.Email
//	simpleUser.Role = u.Role
//	return &simpleUser
//}
