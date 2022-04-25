package catagory

import (
	"capstone/delivery/helper"
	_categoryUseCase "capstone/usecase/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUseCase _categoryUseCase.CategoryUseCaseInterface
}

func NewCategoryHandler(c _categoryUseCase.CategoryUseCaseInterface) CategoryHandler {
	return CategoryHandler{
		categoryUseCase: c,
	}
}

func (uh *CategoryHandler) GetAllCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		catagory, err := uh.categoryUseCase.GetAllCategory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get all catagories"))
		}

		responseCategories := []map[string]interface{}{}
		for i := 0; i < len(catagory); i++ {
			response := map[string]interface{}{
				"id":   catagory[i].ID,
				"name": catagory[i].Name,
			}
			responseCategories = append(responseCategories, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all catagories", responseCategories))
	}
}
