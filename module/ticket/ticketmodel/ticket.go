package ticketmodel

import (
	"errors"
	"food_delivery/common"
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
	Name             string `json:"name" gorm:"column:name;"`
	Email            string `json:"email" gorm:"column:email;"`
	Phone            string `json:"phone" gorm:"column:phone;"`
	JourneyId        string `json:"journey_id" gorm:"column:journey_id;"`
	TrainId          string `json:"train_id" gorm:"column:train_id;"`
	CarriageId       string `json:"carriage_id" gorm:"column:carriage_id;"`
	ChairId          string `json:"chair_id" gorm:"column:chair_id;"`
	JourneyArrive    string `json:"journey_arrive" gorm:"column:journey_arrive;"`
	JourneyDeparture string `json:"journey_departure" gorm:"column:journey_departure;"`
	DepartureTime    int64  `json:"departure_time" gorm:"column:departure_time;"`
	ArriveTime       int64  `json:"arrive_time" gorm:"column:arrive_time;"`
	TrainName        string `json:"train_name" gorm:"column:train_name;"`
	CarriageNumber   int    `json:"carriage_number" gorm:"column:carriage_number;"`
	ChairName        string `json:"chair_name" gorm:"column:chair_name;"`
	ChairType        string `json:"chair_type" gorm:"column:chair_type;"`
}

func (Ticket) TableName() string {
	return "ticker_orders"
}

type TicketCreate struct {
	common.SQLModel  `json:",inline"`
	Name             string `json:"name" gorm:"column:name;"`
	Email            string `json:"email" gorm:"column:email;"`
	Phone            string `json:"phone" gorm:"column:phone;"`
	JourneyId        string `json:"journey_id" gorm:"column:journey_id;"`
	TrainId          string `json:"train_id" gorm:"column:train_id;"`
	CarriageId       string `json:"carriage_id" gorm:"column:carriage_id;"`
	ChairId          string `json:"chair_id" gorm:"column:chair_id;"`
	JourneyArrive    string `json:"journey_arrive" gorm:"column:journey_arrive;"`
	JourneyDeparture string `json:"journey_departure" gorm:"column:journey_departure;"`
	DepartureTime    int64  `json:"departure_time" gorm:"column:departure_time;"`
	ArriveTime       int64  `json:"arrive_time" gorm:"column:arrive_time;"`
	TrainName        string `json:"train_name" gorm:"column:train_name;"`
	CarriageNumber   int    `json:"carriage_number" gorm:"column:carriage_number;"`
	ChairName        string `json:"chair_name" gorm:"column:chair_name;"`
	ChairType        string `json:"chair_type" gorm:"column:chair_type;"`
}

func (TicketCreate) TableName() string {
	return Ticket{}.TableName()
}

// Mask
func (r *Ticket) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeTicketOrder)
}

func (data *TicketCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

// // Validate
// func (data *RestaurantCreate) Validate() error {
// 	data.Name = strings.TrimSpace(data.Name)

// 	if data.Name == "" {
// 		return ErrNameIsEmpty
// 	}

// 	return nil

// }

// Error
var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
