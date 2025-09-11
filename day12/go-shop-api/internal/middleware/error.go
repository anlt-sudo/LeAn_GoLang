package middleware

import (
	"go-shop-api/internal/errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status    int               `json:"status"`
	Message   string            `json:"message"`
	Errors    map[string]string `json:"errors,omitempty"`
	Timestamp time.Time         `json:"timestamp"`
}

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        if len(c.Errors) == 0 {
            return
        }

        err := c.Errors.Last().Err
        if ae, ok := err.(*errors.AppError); ok {
            c.AbortWithStatusJSON(ae.Code, ae)
        } else {
            c.AbortWithStatusJSON(http.StatusInternalServerError, errors.New(http.StatusInternalServerError, err.Error()))
        }
    }
}
