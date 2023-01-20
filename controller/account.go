package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oos/dto"
	"oos/middleware"
	"oos/model"
)

//	@Summary		JWT login
//	@Description	Authenticate to get access token
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	model.Token
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Param			role	path		string	true	"User role (permission or scope)"	Enums(customer, provider)
//	@Router			/account/login/{role} [post]
func Login(c *gin.Context) {
	// HTTP request
	role := c.Param("role")

	// Business logic
	token, err := middleware.CreateAccessToken(role)
	if err != nil {
		dto.Response.
			SetCode(http.StatusInternalServerError).
			SetText(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	result := model.Token{
		UserRole: role,
		JwtToken: token,
	}

	// HTTP response
	dto.Response.
		SetCode(http.StatusOK).
		SetText(http.StatusText(http.StatusOK)).
		SetData(result).
		SendJSON(c)
}
