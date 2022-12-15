package ordermodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"food_delivery/common"
)

type Food struct {
	Id           string        `json:"id" gorm:"column:id"`
	RestaurantId int           `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId   int           `json:"category_id" gorm:"column:category_id;"`
	Name         string        `json:"name" gorm:"column:name;"`
	Images       *common.Image `json:"images" gorm:"column:images;"`
	Price        float64       `json:"price" gorm:"column:price;"`
	Description  string        `json:"description" gorm:"column:description;"`
}

func (f *Food) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Faled to unmarshal JSONB value:", value))
	}

	var food Food
	if err := json.Unmarshal(bytes, &food); err != nil {
		return err
	}

	*f = food

	return nil
}

func (f *Food) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	return json.Marshal(f)
}

type OrderDetail struct {
	common.SQLModel `json:",inline"`
	OrderId         int     `json:"order_id" gorm:"column:order_id;"`
	Discount        float64 `json:"discount" gorm:"column:discount;"`
	Price           float64 `json:"price" gorm:"column:price;"`
	Quantity        int     `json:"quantity" gorm:"column:quantity;"`
	FoodOrigin      *Food   `json:"food_origin" gorm:"column:food_origin;type:json;"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

type OrderDetailCreate struct {
	common.SQLModel `json:",inline"`
	OrderCreateId   int     `json:"order_id" gorm:"column:order_id;"`
	Discount        float64 `json:"discount" gorm:"column:discount;"`
	Price           float64 `json:"price" gorm:"column:price;"`
	Quantity        int     `json:"quantity" gorm:"column:quantity;"`
	FoodOrigin      *Food   `json:"food_origin" gorm:"column:food_origin;type:json;"`
}

func (OrderDetailCreate) TableName() string {
	return "order_details"
}

// Mask
func (r *OrderDetail) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeTicketOrder)
}

func (data *OrderDetailCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
