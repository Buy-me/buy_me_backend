package categorymodel

type Filter struct {
	OwnerId int `json:"owner_id,omitempty" form:"owner_id"`
}
