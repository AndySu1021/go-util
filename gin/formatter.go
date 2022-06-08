package gin

import (
	"github.com/AndySu1021/go-util/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    -1,
		"message": err.Error(),
		"data":    nil,
	})
}

func ErrorAuth(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    -1,
		"message": errors.ErrorAuth.Error(),
		"data":    nil,
	})
}

func ErrorPerm(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"code":    -1,
		"message": errors.ErrorPerm.Error(),
		"data":    nil,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    nil,
	})
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    data,
	})
}

func SuccessWithPagination(c *gin.Context, data interface{}, pagination interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":       0,
		"message":    "Success",
		"data":       data,
		"pagination": pagination,
	})
}
