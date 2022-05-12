package catagory

import (
	"capstone/delivery/helper"
	_entities "capstone/entities"
	_categoryUseCase "capstone/usecase/category"
	"net/http"
	"strconv"

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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed get all catagories", http.StatusBadRequest))
		}

		// responseCategories := []map[string]interface{}{}
		// for i := 0; i < len(catagory); i++ {
		// 	response := map[string]interface{}{
		// 		"id":        catagory[i].ID,
		// 		"name":      catagory[i].Name,
		// 		"icon_name": catagory[i].IconName,
		// 	}
		// 	responseCategories = append(responseCategories, response)
		// }

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all catagories", http.StatusOK, catagory))
	}
}

func (uh *CategoryHandler) CreateCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param _entities.Category

		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		_, err = uh.categoryUseCase.CreateCategory(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("created category failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("created category successfully", http.StatusOK))
	}

}

func (uh *CategoryHandler) UpdateCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		var param _entities.Category
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.Bind(&param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error(), http.StatusBadRequest))
		}
		_, rows, err := uh.categoryUseCase.UpdateCategory(uint(id), param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("updated category failed", http.StatusBadRequest))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("updated category successfully", http.StatusOK))
	}
}

func (uh *CategoryHandler) DeleteCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		err := uh.categoryUseCase.DeleteCategory(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("deleted category failed", http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("deleted category successfully", http.StatusOK))
	}
}
