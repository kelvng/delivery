package restaurantmodel

type Filter struct {
	OwnerID int   `json:"owner_id,omitempty" form:"owner_id"`
	Status  []int `json:"-"`
}
