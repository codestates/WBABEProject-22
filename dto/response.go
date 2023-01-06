package dto

import "github.com/gin-gonic/gin"

var Resp Response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response Response) SetCode(status int) Response {
	response.Code = status
	return response
}

func (response Response) SetMessage(message string) Response {
	response.Message = message
	return response
}

func (response Response) SetData(data interface{}) Response {
	response.Data = data
	return response
}

func (response Response) SendJSON(c *gin.Context) {
	c.JSON(response.Code, response)
}

func (response Response) SendIndentedJSON(c *gin.Context) {
	c.IndentedJSON(response.Code, response)
}

func (response Response) AbortWithStatusJSON(c *gin.Context) {
	c.AbortWithStatusJSON(response.Code, response)
}

// References
// https://github.com/Shipu/artifact/blob/master/response.go
