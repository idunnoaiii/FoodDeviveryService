package restaurantstorage

import (
	"200lab/common"
	"200lab/modules/restaurant/restaurantmodel"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {

	db := s.db

	var data restaurantmodel.Restaurant

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := s.db.Where(conditions).First(&data).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
