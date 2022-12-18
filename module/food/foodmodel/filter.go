package foodmodel

type Filter struct {
	OwnerId    int `json:"owner_id,omitempty" form:"owner_id"`
	CategoryId int `json:"category_id,omitempty" form:"category_id"`
}
