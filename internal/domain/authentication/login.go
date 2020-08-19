package authentication

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthy-buddy/internal/model"
	"github.com/jmramos02/healthy-buddy/internal/utility"
)

type LoginRequest struct {
	Email    string `json:"email_address"`
	Password string `json:"password"`
}

func Login(ctx context.Context, request LoginRequest) (response AuthenticationResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	var user model.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return response, &utility.HttpError{
			StatusCode: 401,
			Message:    "UNAUTHORIZED",
		}
	}

	if err := utility.CompareToHash(user.Password, request.Password); err != nil {
		return response, &utility.HttpError{
			StatusCode: 401,
			Message:    "UNAUTHORIZED",
		}
	}

	if user.Type == "customer" {
		var customer model.Customer
		if err := db.Where("user_id = ?", user.ID).First(&customer).Error; err != nil {
			return response, &utility.HttpError{
				StatusCode: 500,
				Message:    "SOMETHING_WENT_WRONG",
			}
		}

		response.Data.AcessToken = encodeUserInfo(user, customer)
		return response, nil
	}

	if user.Type == "dietitian" {
		var dietitian model.Dietitian
		if err := db.Where("user_id = ?", user.ID).First(&dietitian).Error; err != nil {
			return response, &utility.HttpError{
				StatusCode: 500,
				Message:    "SOMETHING_WENT_WRONG",
			}
		}

		response.Data.AcessToken = encodeUserInfo(user, dietitian)
		return response, nil
	}

	return response, &utility.HttpError{
		StatusCode: 500,
		Message:    "SOMETHING_WENT_WRONG",
	}
}
