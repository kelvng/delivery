package userstore

import (
	"awesomeProject1/common"
	usermodel "awesomeProject1/module/user/model"
	"context"
)

func (s *sqlStore) CreateUser(context context.Context, data *usermodel.UserCreate) error {

	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
