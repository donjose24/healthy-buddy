package authentication

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthy-buddy/internal/model"
	"github.com/jmramos02/healthy-buddy/internal/utility"
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
	Specialty         string  `json:"specialty"`
	YearsOfExperience int     `json:"years_of_experience"`
}

type AuthenticationResponse struct {
	Data AccessToken `json:"data"`
}

type AccessToken struct {
	AcessToken string `json:"access_token"`
}

func Register(ctx context.Context, request RegisterRequest, userType string) (response AuthenticationResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	var user model.User

	if err = db.Where("email = ?", request.Email).First(&user).Error; err == nil {
		return response, &utility.HttpError{
			StatusCode: 400,
			Message:    "INVALID_PARAMETERS",
			Errors: map[string]string{
				"email": "already_exists",
			},
		}

	}

	fmt.Println(err == nil)

	user = model.User{
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

	switch userType {
	case "customer":
		{
			var userData model.Customer
			userData = model.Customer{
				Weight:            request.Weight,
				Height:            request.Height,
				Gender:            request.Gender,
				Goal:              request.Goal,
				Allergy:           request.Allergy,
				DietaryPreference: request.DietaryPreference,
				UserID:            user.ID,
			}

			if err = db.Create(&userData).Error; err != nil {
				return response, &utility.HttpError{
					StatusCode: 500,
					Message:    "SOMETHING_WENT_WRONG",
				}
			}

			response.Data.AcessToken = encodeUserInfo(user, userData)
			break
		}

	case "dietitian":
		{
			dietitianData := model.Dietitian{
				YearsOfExperience: request.YearsOfExperience,
				Specialty:         request.Specialty,
			}

			if err = db.Create(&dietitianData).Error; err != nil {
				return response, &utility.HttpError{
					StatusCode: 500,
					Message:    "SOMETHING_WENT_WRONG",
				}
			}

			response.Data.AcessToken = encodeUserInfo(user, dietitianData)
			break
		}
	}

	return response, nil
}
