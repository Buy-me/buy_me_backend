package ordermodel

type OrderUpdate struct {
	State string `json:"tracking_state" gorm:"column:state;"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
