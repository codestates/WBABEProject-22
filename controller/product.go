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
//	@Param			newProduct	body		dto.ProductCreate	true	"A new product to add"
//	@Success		200			{object}	model.Product
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Router			/provider/products [post]
func CreateProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	var product dto.ProductCreate
	err := c.BindJSON(&product)
	if err != nil {
		dto.Resp.
		SetCode(http.StatusBadRequest).
		SetMessage(http.StatusText(http.StatusBadRequest)).
		SetData(err.Error()).
		AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.CreateProduct(ctx, product)
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

//	@Summary		Get all products
//	@Description	List all products available to customers
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Product
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/customer/products [get]
func GetProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Business logic
	result, err := service.GetProducts(ctx)
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

//	@Summary		Get a product
//	@Description	Show a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"The product to show"
//	@Success		200		{object}	model.Product
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/products/{code} [get]
func GetProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	// Business logic
	result, err := service.GetProduct(ctx, productCode)
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

//	@Summary		Update a product
//	@Description	Modify an existing product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string			true	"Product code"
//	@Param			product	body		model.Product	true	"The product to modify"
//	@Success		200		{object}	model.Product
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/provider/products/{code} [put]
func UpdateProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	var product dto.ProductUpdate
	err := c.BindJSON(&product)
	if err != nil {
		panic(err)
	}

	// Business logic
	result, err := service.UpdateProduct(ctx, productCode, product)
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
func DeleteProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	productCode := c.Param("code")

	// Business logic
	result, err := service.DeleteProduct(ctx, productCode)
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
