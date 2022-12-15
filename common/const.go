package common

const (
	DbTypeRestaurant  = 1
	DbTypeFood        = 2
	DbTypeCategory    = 3
	DbTypeUser        = 4
	DbTypeTicketOrder = 5
	DbTypeOrder       = 6
)

const CurrentUser = "user"

const (
	TopicUserLikeRestaurant   = "TopicUserLikeRestaurant"
	TopicUserUnLikeRestaurant = "TopicUserUnLikeRestaurant"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
