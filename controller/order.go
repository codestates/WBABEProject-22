package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"oos/dto"
	"oos/service"
)

//	@Summary		Create a new order
//	@Description	Add an order document to the orders collection
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		dto.OrderCreate	true	"A new order to submit"
//	@Success		200		{object}	model.Order
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/orders [post]
//	@Security		ApiKeyAuth
func CreateOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	var order dto.OrderCreate
	err := c.BindJSON(&order)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.CreateOrder(ctx, order)
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

//	@Summary		List all orders
//	@Description	Show all orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Order
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/provider/orders [get]
//	@Security		ApiKeyAuth
func ListOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Business logic
	result, err := service.ListOrders(ctx)
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

//	@Summary		List all active orders
//	@Description	Show all orders currently active by username
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}		model.Order
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Param			username	path		string	true	"Username"
//	@Router			/customer/{username}/orders/active [get]
//	@Security		ApiKeyAuth
func ListOrdersActive(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	username := c.Param("username")

	// Business logic
	result, err := service.ListOrdersActive(ctx, username)
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

//	@Summary		List all past orders
//	@Description	Show all order history by username
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}		model.Order
//	@Failure		400			{object}	error
//	@Failure		404			{object}	error
//	@Failure		500			{object}	error
//	@Param			username	path		string	true	"Username"
//	@Router			/customer/{username}/orders/history [get]
//	@Security		ApiKeyAuth
func ListOrdersHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	username := c.Param("username")

	// Business logic
	result, err := service.ListOrdersHistory(ctx, username)
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

//	@Summary		Get an order
//	@Description	Show an order by order ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Order
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Param			id	path		string	true	"Order ID"
//	@Router			/customer/orders/{id} [get]
//	@Security		ApiKeyAuth
func GetOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderID := c.Param("id")

	// Business logic
	result, err := service.GetOrder(ctx, orderID)
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

//	@Summary		Get order status
//	@Description	Show the current status of an order by order ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Param			id	path		string	true	"Order ID"
//	@Router			/customer/orders/{id}/status [get]
//	@Security		ApiKeyAuth
func GetOrderStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderID := c.Param("id")

	// Business logic
	result, err := service.GetOrderStatus(ctx, orderID)
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

//	@Summary		Update order status
//	@Description	Modify order status
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Order ID"
//	@Param			order	body		dto.OrderUpdateStatus	true	"Updated order status"
//	@Success		200		{object}	model.Order
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/provider/orders/{id}/status [put]
//	@Security		ApiKeyAuth
func UpdateOrderStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderID := c.Param("id")

	var order dto.OrderUpdateStatus
	err := c.BindJSON(&order)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.UpdateOrderStatus(ctx, orderID, order)
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

//	@Summary		Update order items
//	@Description	Modify order items
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"Order ID"
//	@Param			order	body		dto.OrderUpdateCart	true	"New items to order"
//	@Success		200		{object}	model.Order
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/orders/{id}/cart [put]
//	@Security		ApiKeyAuth
func UpdateOrderItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderID := c.Param("id")

	var order dto.OrderUpdateCart
	err := c.BindJSON(&order)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.UpdateOrderItems(ctx, orderID, order)
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

//	@Summary		Delete order items
//	@Description	Remove order items
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Order ID"
//	@Param			order	body		[]string	true	"Items to delete"
//	@Success		200		{object}	model.Order
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/customer/orders/{id}/cart [delete]
//	@Security		ApiKeyAuth
func DeleteOrderItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderID := c.Param("id")

	var products []string
	err := c.BindJSON(&products)
	if err != nil {
		dto.Response.
			SetCode(http.StatusBadRequest).
			SetText(http.StatusText(http.StatusBadRequest)).
			SetData(err.Error()).
			AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.DeleteOrderItems(ctx, orderID, products)
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
