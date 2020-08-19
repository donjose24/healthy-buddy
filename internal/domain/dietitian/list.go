package dietitian

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthy-buddy/internal/model"
	"github.com/jmramos02/healthy-buddy/internal/utility"
)

type ListDietitianResponse struct {
	Data []model.Dietitian `json:"data"`
}

func List(ctx context.Context) (response ListDietitianResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	var dietitians []model.Dietitian

	if err := db.Preload("User").Find(&dietitians).Error; err != nil {
		return response, &utility.HttpError{
			StatusCode: 500,
			Message:    "SOMETHING_WENT_WRONG",
		}
	}

	response.Data = dietitians
	return response, nil
}
