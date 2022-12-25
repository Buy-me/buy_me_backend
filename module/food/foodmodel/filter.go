package foodmodel

type Filter struct {
	OwnerId    int    `json:"owner_id,omitempty" form:"owner_id"`
	CategoryId int    `json:"category_id,omitempty" form:"category_id"`
	Sort       string `json:"sort,omitempty" form:"sort"`
	MinPrice   int    `json:"min_price,omitempty" form:"min_price"`
	MaxPrice   int    `json:"max_price,omitempty" form:"max_price"`
	Search     string `json:"search,omitempty" form:"search"`
	Rating     int    `json:"rating,omitempty" form:"rating"`
}
