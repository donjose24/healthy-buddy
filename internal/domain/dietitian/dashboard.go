package dietitian

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthy-buddy/internal/model"
	"github.com/jmramos02/healthy-buddy/internal/utility"
)

type GetDietitianDashboardResponse struct {
	Data model.Dietitian `json:"data"`
}

func Fetch(ctx context.Context) (response GetDietitianDashboardResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	user := ctx.Value("user").(model.User)

	var dietitian model.Dietitian

	if err := db.Preload("Customers").Preload("Customers.User").Preload("User").Where("user_id = ?", user.ID).First(&dietitian).Error; err != nil {
		return response, &utility.HttpError{
			StatusCode: 401,
			Message:    "UNAUTHORIZED",
		}
	}

	response.Data = dietitian
	return response, nil
}
