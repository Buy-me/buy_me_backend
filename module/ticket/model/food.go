package ticketmodel

import (
	"errors"
	"food_delivery/common"
	"time"
)

type RestaurantType string

const EntityName = "Ticket"

/*
 `id` int NOT NULL AUTO_INCREMENT,
  `journey_id` varchar(50) NOT NULL,
  `journey_arrive` varchar(50) NOT NULL,
  `journey_departure` varchar(50) NOT NULL,
  `departure_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `arrive_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `train_id` int,
  `train_name` varchar(50) NOT NULL,
  `carriage_id` int,
  `carriage_number` int,
  `chair_id` varchar(50) NOT NULL,
  `chair_name` varchar(50) NOT NULL,
  `chair_type` varchar(50) NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
*/

type Ticket struct {
	common.SQLModel  `json:",inline"`
	JourneyId        int        `json:"journey_id" gorm:"column:journey_id;"`
	TrainId          int        `json:"train_id" gorm:"column:train_id;"`
	CarriageId       int        `json:"carriage_id" gorm:"column:carriage_id;"`
	ChairId          int        `json:"chair_id" gorm:"column:chair_id;"`
	JourneyArrive    string     `json:"journey_arrive" gorm:"column:journey_arrive;"`
	JourneyDeparture string     `json:"journey_departure" gorm:"column:journey_departure;"`
	DepartureTime    *time.Time `json:"departure_time" gorm:"column:departure_time;"`
	ArriveTime       *time.Time `json:"arrive_time" gorm:"column:arrive_time;"`
	TrainName        string     `json:"train_name" gorm:"column:train_name;"`
	CarriageNumber   int        `json:"carriage_number" gorm:"column:carriage_number;"`
	ChairName        string     `json:"chair_name" gorm:"column:chair_name;"`
	ChairType        string     `json:"chair_type" gorm:"column:chair_type;"`
}

func (Ticket) TableName() string {
	return "ticker_orders"
}

// Mask
func (r *Ticket) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeTicketOrder)
}

// type Restaurant struct {
// 	common.SQLModel `json:",inline"`
// 	Name            string             `json:"name" gorm:"column:name;"`
// 	Addr            string             `json:"addr" gorm:"column:addr;"`
// 	Type            RestaurantType     `json:"type" gorm:"column:type;"`
// 	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
// 	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
// 	UserId          int                `json:"-" gorm:"column:user_id;"`
// 	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
// 	LikedCount      int                `json:"liked_count" gorm:"column:liked_count;"`
// }

// type FoodCreate struct {
// 	common.SQLModel `json:",inline"`
// 	RestaurantId    int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
// 	Name            string        `json:"name" gorm:"column:name;"`
// 	Images          *common.Image `json:"images" gorm:"column:images;"`
// 	Price           float64       `json:"price" gorm:"column:price;"`
// 	Description     string        `json:"description" gorm:"column:description;"`
// }

// func (FoodCreate) TableName() string {
// 	return Food{}.TableName()
// }

// type FoodUpdate struct {
// 	Name        string        `json:"name" gorm:"column:name;"`
// 	Images      *common.Image `json:"images" gorm:"column:images;"`
// 	Price       float64       `json:"price" gorm:"column:price;"`
// 	Description string        `json:"description" gorm:"column:description;"`
// }

// func (FoodUpdate) TableName() string {
// 	return Food{}.TableName()
// }

// // Validate
// func (data *FoodCreate) Validate() error {
// 	data.Name = strings.TrimSpace(data.Name)

// 	if data.Name == "" {
// 		return ErrNameIsEmpty
// 	}

// 	return nil
// }

// // Mask
// func (r *Food) Mask(isAdminOrOwner bool) {
// 	r.GenUID(common.DbTypeRestaurant)
// }

// func (data *FoodCreate) Mask(isAdminOrOwner bool) {
// 	data.GenUID(common.DbTypeRestaurant)
// }

// type RestaurantCreate struct {
// 	common.SQLModel `json:",inline"`
// 	Name            string         `json:"name" gorm:"column:name;"`
// 	UserId          int            `json:"-" gorm:"column:user_id;"`
// 	Addr            string         `json:"addr" gorm:"column:addr;"`
// 	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
// 	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
// }

// func (RestaurantCreate) TableName() string {
// 	return Restaurant{}.TableName()
// }

// type RestaurantUpdate struct {
// 	Name  *string        `json:"name" gorm:"column:name;"`
// 	Addr  *string        `json:"addr" gorm:"column:addr;"`
// 	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
// 	Cover *common.Images `json:"cover" gorm:"column:cover;"`
// }

// func (RestaurantUpdate) TableName() string {
// 	return Restaurant{}.TableName()
// }

// // Validate
// func (data *RestaurantCreate) Validate() error {
// 	data.Name = strings.TrimSpace(data.Name)

// 	if data.Name == "" {
// 		return ErrNameIsEmpty
// 	}

// 	return nil

// }

// Mask
// func (r *Food) Mask(isAdminOrOwner bool) {
// 	r.GenUID(common.DbTypeRestaurant)

// 	if u := r.; u != nil {
// 		u.Mask(isAdminOrOwner)
// 	}
// }

// func (data *FoodCreate) Mask(isAdminOrOwner bool) {
// 	data.GenUID(common.DbTypeRestaurant)
// }

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
