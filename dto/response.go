package dto

import "github.com/gin-gonic/gin"

var Response HTTPResponse

type HTTPResponse struct {
	Code int         `json:"code"`
	Text string      `json:"text"`
	Data interface{} `json:"data"`
}

func (response HTTPResponse) SetCode(statusCode int) HTTPResponse {
	response.Code = statusCode
	return response
}

func (response HTTPResponse) SetText(statusMessage string) HTTPResponse {
	response.Text = statusMessage
	return response
}

func (response HTTPResponse) SetData(rawResponse interface{}) HTTPResponse {
	response.Data = rawResponse
	return response
}

func (response HTTPResponse) SendJSON(c *gin.Context) {
	c.JSON(response.Code, response)
}

func (response HTTPResponse) SendIndentedJSON(c *gin.Context) {
	c.IndentedJSON(response.Code, response)
}

func (response HTTPResponse) AbortWithStatusJSON(c *gin.Context) {
	c.AbortWithStatusJSON(response.Code, response)
}

// References
// https://github.com/Shipu/artifact/blob/master/response.go
