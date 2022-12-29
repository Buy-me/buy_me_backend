package usermodel

type Filter struct {
	Sort   string `json:"sort,omitempty" form:"sort"`
	Search string `json:"search,omitempty" form:"search"`
}
