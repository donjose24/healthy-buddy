package customer

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthy-buddy/internal/model"
	"github.com/jmramos02/healthy-buddy/internal/utility"
)

type GetCustomerDashboardResponse struct {
	Data model.Customer `json:"data"`
}

func Fetch(ctx context.Context) (response GetCustomerDashboardResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	user := ctx.Value("user").(model.User)

	var customer model.Customer

	if err := db.Preload("Dietitian").Preload("MealPlan").Preload("MealPlan.MealPlanEntries").Preload("User").Preload("Dietitian.User").Where("user_id = ?", user.ID).First(&customer).Error; err != nil {
		return response, &utility.HttpError{
			StatusCode: 401,
			Message:    "UNAUTHORIZED",
		}
	}

	response.Data = customer
	return response, nil
}
