package ticketmodel

import (
	"food_delivery/common"
)

const NameTicketDetail = "TicketDetail"

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

type TicketDetail struct {
	common.SQLModel `json:",inline"`
	TicketId        int     `json:"order_id" gorm:"column:order_id;autoIncrement:false;"`
	Name            string  `json:"name" gorm:"column:name;"`
	Age             string  `json:"age" gorm:"column:age;"`
	National        string  `json:"national" gorm:"column:national;"`
	Gender          string  `json:"gender" gorm:"column:gender;"`
	Price           float64 `json:"price" gorm:"column:price;"`
}

func (TicketDetail) TableName() string {
	return "ticket_order_detail"
}

type TicketDetailCreate struct {
	common.SQLModel `json:",inline"`
	TicketCreateId  int     `json:"order_id" gorm:"column:order_id;autoIncrement:false;"`
	Name            string  `json:"name" gorm:"column:name;"`
	Age             string  `json:"age" gorm:"column:age;"`
	National        string  `json:"national" gorm:"column:national;"`
	Gender          string  `json:"gender" gorm:"column:gender;"`
	Price           float64 `json:"price" gorm:"column:price;"`
}

func (TicketDetailCreate) TableName() string {
	return "ticket_order_detail"
}

// Mask
func (r *TicketDetail) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeTicketOrder)
}

func (data *TicketDetailCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
