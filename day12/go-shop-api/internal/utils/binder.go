package utils

import (
	"net/http"

	appErr "go-shop-api/internal/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		fieldErrors := make(map[string]string)
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				fieldErrors[fe.Field()] = fe.Error()
			}
		} else {
			fieldErrors["body"] = err.Error()
		}
		return appErr.NewWithErrors(http.StatusBadRequest, "Invalid request body", fieldErrors)
	}
	return nil
}
