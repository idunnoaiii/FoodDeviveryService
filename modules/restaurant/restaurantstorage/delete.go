package restaurantstorage

import (
	"200lab/common"
	"200lab/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where(id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
