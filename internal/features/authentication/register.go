package authentication

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthybuddy/internal/model"
	"github.com/jmramos02/healthybuddy/internal/utility"
)

type RegisterRequest struct {
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Email             string  `json:"email_address"`
	Password          string  `json:"password"`
	Goal              string  `json:"goal"`
	Allergy           string  `json:"allergy"`
	Weight            float64 `json:"weight"`
	Height            string  `json:"height"`
	DietaryPreference string  `json:"dietary_preference"`
	Gender            string  `json:"gender"`
}

type AuthenticationResponse struct {
	Data AccessToken `json:"data"`
}

type AccessToken struct {
	AcessToken string `json:"access_token"`
}

func RegisterCustomer(ctx context.Context, request RegisterRequest, userType string) (response AuthenticationResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	user := model.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  utility.HashString(request.Password),
		Type:      userType,
	}

	if err = db.Create(&user).Error; err != nil {
		return response, &utility.HttpError{
			StatusCode: 500,
			Message:    "SOMETHING_WENT_WRONG",
		}
	}

	var userData model.Customer
	switch userType {
	case "customer":
		{
			userData = model.Customer{
				Weight:            request.Weight,
				Height:            request.Height,
				Gender:            request.Gender,
				Goal:              request.Goal,
				Allergy:           request.Allergy,
				DietaryPreference: request.DietaryPreference,
				UserID:            user.ID,
			}
			break
		}
	}

	if err = db.Create(&userData).Error; err != nil {
		return response, &utility.HttpError{
			StatusCode: 500,
			Message:    "SOMETHING_WENT_WRONG",
		}
	}

	response.Data.AcessToken = encodeUserInfo(user, userData)

	return response, nil
}
