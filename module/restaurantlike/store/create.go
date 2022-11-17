package restaurantlikestorage

import (
	"awesomeProject1/common"
	restaurantlikemodel "awesomeProject1/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	
	return nil
}
