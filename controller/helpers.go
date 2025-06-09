package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ReadJson(c *gin.Context, data any) error {
	err := c.BindJSON(data)
	if err != nil {
		fmt.Println("error during read the input body", err)
		return err
	}

	return nil
}

func ErrorJson(c *gin.Context, statusCode int, ErrorMessage error) {
	var ErrorJsonData struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	ErrorJsonData.Error = true
	ErrorJsonData.Message = ErrorMessage.Error()
	c.JSON(statusCode, ErrorJsonData)

}

func WriteJson(c *gin.Context, statusCode int, message string, returnData any) {
	var Response struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
	}

	Response.Data = returnData
	Response.Message = message
	c.JSON(statusCode, Response)

}
