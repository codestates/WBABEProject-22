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
func CreateOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	var order dto.OrderCreate
	err := c.BindJSON(&order)
	if err != nil {
		dto.Resp.
		SetCode(http.StatusBadRequest).
		SetMessage(http.StatusText(http.StatusBadRequest)).
		SetData(err.Error()).
		AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.CreateOrder(ctx, order)
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

//	@Summary		Get all orders
//	@Description	List all orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Order
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/provider/orders [get]
func GetOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Business logic
	result, err := service.GetOrders(ctx)
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

//	@Summary		Get all active orders
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
func GetOrdersActive(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	username := c.Param("username")

	// Business logic
	result, err := service.GetOrdersActive(ctx, username)
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

//	@Summary		Get all past orders
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
func GetOrdersHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	username := c.Param("username")

	// Business logic
	result, err := service.GetOrdersHistory(ctx, username)
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
func GetOrder(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderId := c.Param("id")

	// Business logic
	result, err := service.GetOrder(ctx, orderId)
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
func GetOrderStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderId := c.Param("id")

	// Business logic
	result, err := service.GetOrderStatus(ctx, orderId)
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
func UpdateOrderStatus(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderId := c.Param("id")

	var order dto.OrderUpdateStatus
	err := c.BindJSON(&order)
	if err != nil {
		dto.Resp.
		SetCode(http.StatusBadRequest).
		SetMessage(http.StatusText(http.StatusBadRequest)).
		SetData(err.Error()).
		AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.UpdateOrderStatus(ctx, orderId, order)
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
func UpdateOrderItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// HTTP request
	orderId := c.Param("id")

	var order dto.OrderUpdateCart
	err := c.BindJSON(&order)
	if err != nil {
		dto.Resp.
		SetCode(http.StatusBadRequest).
		SetMessage(http.StatusText(http.StatusBadRequest)).
		SetData(err.Error()).
		AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.UpdateOrderItems(ctx, orderId, order)
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
func DeleteOrderItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// HTTP request
	orderId := c.Param("id")

	var products []string
	err := c.BindJSON(&products)
	if err != nil {
		dto.Resp.
		SetCode(http.StatusBadRequest).
		SetMessage(http.StatusText(http.StatusBadRequest)).
		SetData(err.Error()).
		AbortWithStatusJSON(c)
		return
	}

	// Business logic
	result, err := service.DeleteOrderItems(ctx, orderId, products)
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

/* [코드리뷰]
 * request 요청에서 발생하는 timeout 처리를 잘 해주셨습니다.
 * request 요청에 대한 다양한 상황을 미리 예측하여 방어적으로 코드를 구현해주시는 방식은
 * 보다 견고한 코드를 만들 수 있는 좋은 방법입니다.
 * 들어오는 api request에 대해서 user validation 코드를 넣어보시는 건 어떨까요?
 * server에서는 항상 client를 의심하는 방어적 코딩 스타일을 수행해야 합니다.
 * 주문자와, 피주문자의 타입을 구분하여 주문자가 피주문자의 api를 요청할 수 없게끔 만들어보시는 방법도 추천드립니다.
 */