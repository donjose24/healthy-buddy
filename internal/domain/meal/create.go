package meal

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jmramos02/healthy-buddy/internal/model"
	"github.com/jmramos02/healthy-buddy/internal/utility"
)

type CreateMealPlanRequest struct {
	StartDate       string     `json:"start_date"`
	EndDate         string     `json:"end_date"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	Remarks         string     `json:"remarks"`
	CustomerID      uint       `json:"customer"`
	MealPlanEntrees []MealPlan `json:"meals"`
}

type MealPlan struct {
	FoodName    string  `json:"food_name"`
	Protein     float64 `json:"protein"`
	Fat         float64 `json:"fat"`
	Carb        float64 `json:"carb"`
	Calories    float64 `json:"calories"`
	Grams       float64 `json:"grams"`
	Description string  `json:"description"`
	MealTime    string  `json:"break_fast"`
	Date        string  `json:"date"`
}

type MealPlanResponse struct {
	Data model.MealPlan `json:"data"`
}

func CreateMealPlan(ctx context.Context, request CreateMealPlanRequest) (response MealPlanResponse, err error) {
	db := ctx.Value("db").(*gorm.DB)
	dietitian := ctx.Value("details").(map[string]interface{})
	ID := uint(dietitian["id"].(float64))
	fmt.Println(request)

	mealPlan := model.MealPlan{
		CustomerID:  request.CustomerID,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Title:       request.Title,
		Remarks:     request.Remarks,
		Description: request.Description,
		DietitianID: ID,
	}

	err = db.Create(&mealPlan).Error

	if err != nil {
		return response, &utility.HttpError{
			StatusCode: 500,
			Message:    "SOMETHING_WENT_WRONG",
		}
	}

	for _, entry := range request.MealPlanEntrees {
		mealPlanEntry := model.MealPlanEntry{
			Calories:    entry.Calories,
			Carb:        entry.Carb,
			Date:        entry.Date,
			Protein:     entry.Protein,
			Food:        entry.FoodName,
			Description: entry.Description,
			MealTime:    entry.MealTime,
			Grams:       entry.Grams,
			Fat:         entry.Fat,
			MealPlanID:  mealPlan.ID,
		}

		err = db.Create(&mealPlanEntry).Error

		if err != nil {
			return response, &utility.HttpError{
				StatusCode: 500,
				Message:    "SOMETHING_WENT_WRONG",
			}
		}
	}

	db.Preload("MealPlanEntries").Where("id = ?", mealPlan.ID).First(&mealPlan)

	response.Data = mealPlan
	return response, nil
}
