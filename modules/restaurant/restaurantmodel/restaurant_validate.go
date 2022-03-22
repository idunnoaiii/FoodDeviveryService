package restaurantmodel

import (
	"errors"
	"strings"
)

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
