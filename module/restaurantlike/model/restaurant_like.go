package restaurantlikemodel

import (
	"awesomeProject1/common"
	"fmt"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this restaurant"),
		fmt.Sprintf("ErrCannotDislikeRestaurant"),
	)
}
