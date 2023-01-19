package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"oos/dto"
	"oos/service"
)

//	@Summary		Create a new product
//	@Description	Add a product document to the products collection
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		dto.ProductCreate	true	"A new product to add"
//	@Success		200		{object}	model.Product
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/provider/products [post]
//	@Security		ApiKeyAuth
func CreateProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	var product dto.ProductCreate
	err := c.BindJSON(&product)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.CreateProduct(ctx, product)
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

//	@Summary		List all products
//	@Description	Show all products available to customers
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			sort	query		string	true	"Parameter used to sort products"	Enums(ratings, reorders, likes, time)
//	@Success		200		{array}		model.ProductView
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/products [get]
//	@Security		ApiKeyAuth
func ListProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	sortBy := c.Query("sort")

	// Business logic
	result, err := service.ListProducts(ctx, sortBy)
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

//	@Summary		Get a product
//	@Description	Show a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"The product to show"
//	@Success		200		{object}	model.ProductView
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/products/{code} [get]
//	@Security		ApiKeyAuth
func GetProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	// Business logic
	result, err := service.GetProduct(ctx, productCode)
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

//	@Summary		Update a product
//	@Description	Modify an existing product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string				true	"Product code"
//	@Param			product	body		dto.ProductUpdate	true	"The product to modify"
//	@Success		200		{object}	model.Product
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/provider/products/{code} [put]
//	@Security		ApiKeyAuth
func UpdateProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	var product dto.ProductUpdate
	err := c.BindJSON(&product)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.UpdateProduct(ctx, productCode, product)
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

//	@Summary		Delete a product
//	@Description	Remove an existing product: toggle canView flag to false
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"The product to delete"
//	@Success		200		{object}	model.Product
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/provider/products/{code} [delete]
//	@Security		ApiKeyAuth
func DeleteProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	// Business logic
	result, err := service.DeleteProduct(ctx, productCode)
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
