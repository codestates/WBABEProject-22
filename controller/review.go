package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"oos/dto"
	"oos/service"
)

//	@Summary		Create a new review
//	@Description	Add a review document to the reviews collection
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Order ID"
//	@Param			review	body		dto.ReviewOrderCreate	true	"A new review to add"
//	@Success		200		{object}	model.ReviewOrder
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/reviews/orders/{id} [post]
func CreateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderId := c.Param("id")

	var review dto.ReviewOrderCreate
	err := c.BindJSON(&review)
	if err != nil {
		panic(err)
	}

	// Business logic
	result, err := service.CreateReview(ctx, orderId, review)
	if err != nil {
		dto.Resp.
			SetCode(http.StatusInternalServerError).
			SetMessage(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	// HTTP response
	dto.Resp.
		SetCode(http.StatusCreated).
		SetMessage(http.StatusText(http.StatusCreated)).
		SetData(result).
		SendJSON(c)
}

//	@Summary		Get all reviews
//	@Description	List all reviews
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.ReviewOrder
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/provider/reviews/orders [get]
func GetReviews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Business logic
	result, err := service.GetReviews(ctx)
	if err != nil {
		dto.Resp.
			SetCode(http.StatusInternalServerError).
			SetMessage(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	// HTTP response
	dto.Resp.
		SetCode(http.StatusOK).
		SetMessage(http.StatusText(http.StatusOK)).
		SetData(result).
		SendJSON(c)
}

//	@Summary		Get a review of a product
//	@Description	Show a review of a product
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"Product code"
//	@Success		200		{array}		model.ReviewProduct
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/reviews/products/{code} [get]
func GetReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	// Business logic
	result, err := service.GetReview(ctx, productCode)
	if err != nil {
		dto.Resp.
			SetCode(http.StatusInternalServerError).
			SetMessage(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	// HTTP response
	dto.Resp.
		SetCode(http.StatusOK).
		SetMessage(http.StatusText(http.StatusOK)).
		SetData(result).
		SendJSON(c)
}
