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

/* [코드리뷰]
 * Response 모델을 잘 구현해주셔서, 
 * 다양한 controller에서 동작하는 코드들의 return 방식을 정확하게 통일시켜주었습니다.
 * 해당 요청의 응답을 받는 client 입장에서 응답을 처리하기에 보다 정확하고 견고한 방식의 코드라고 생각됩니다.
 * 정말 잘 만들어주셨습니다.
 */

// References
// https://github.com/Shipu/artifact/blob/master/response.go
