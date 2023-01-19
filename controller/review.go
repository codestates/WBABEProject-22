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
//	@Security		ApiKeyAuth
func CreateReview(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderID := c.Param("id")

	var review dto.ReviewOrderCreate
	err := c.BindJSON(&review)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.CreateReview(ctx, orderID, review)
	if err != nil {
		dto.Response.
			SetCode(http.StatusInternalServerError).
			SetText(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	// HTTP response
	dto.Response.
		SetCode(http.StatusCreated).
		SetText(http.StatusText(http.StatusCreated)).
		SetData(result).
		SendJSON(c)
}

//	@Summary		List all reviews
//	@Description	Show all reviews
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.ReviewOrder
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/provider/reviews/orders [get]
//	@Security		ApiKeyAuth
func ListReviews(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Business logic
	result, err := service.ListReviews(ctx)
	if err != nil {
		dto.Response.
			SetCode(http.StatusInternalServerError).
			SetText(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	// HTTP response
	dto.Response.
		SetCode(http.StatusOK).
		SetText(http.StatusText(http.StatusOK)).
		SetData(result).
		SendJSON(c)
}

//	@Summary		List all reviews of a product
//	@Description	Show all reviews of a product
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"Product code"
//	@Success		200		{array}		model.ReviewProduct
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/reviews/products/{code} [get]
//	@Security		ApiKeyAuth
func ListReviewsProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	// Business logic
	result, err := service.ListReviewsProduct(ctx, productCode)
	if err != nil {
		dto.Response.
			SetCode(http.StatusInternalServerError).
			SetText(http.StatusText(http.StatusInternalServerError)).
			SetData(err.Error()).
			SendJSON(c)
		return
	}

	// HTTP response
	dto.Response.
		SetCode(http.StatusOK).
		SetText(http.StatusText(http.StatusOK)).
		SetData(result).
		SendJSON(c)
}
