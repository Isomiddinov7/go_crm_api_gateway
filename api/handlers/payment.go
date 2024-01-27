package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreatePayment godoc
// @ID create_payment
// @Router /payment [POST]
// @Summary Create Payment
// @Description  Create Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param profile body users_service.CreatePayment true "CreatePaymentBody"
// @Success 200 {object} http.Response{data=users_service.Payment} "GetPaymentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePayment(c *gin.Context) {

	var payment users_service.CreatePayment

	err := c.ShouldBindJSON(&payment)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PaymentService().Create(
		c.Request.Context(),
		&payment,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetPaymentByID godoc
// @ID get_payment_by_id
// @Router /payment/{id} [GET]
// @Summary Get Payment  By ID
// @Description Get Payment  By ID
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Payment} "PaymentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPaymentByID(c *gin.Context) {

	paymentID := c.Param("id")

	if !util.IsValidUUID(paymentID) {
		h.handleResponse(c, http.InvalidArgument, "payment id is an invalid uuid")
		return
	}

	resp, err := h.services.PaymentService().GetById(
		context.Background(),
		&users_service.PaymentPrimaryKey{
			Id: paymentID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetPaymentList godoc
// @ID get_payment_list
// @Router /payment [GET]
// @Summary Get Payment s List
// @Description  Get Payment s List
// @Tags Payment
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListPaymentResponse} "GetAllPaymentResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPaymentList(c *gin.Context) {

	if c.GetHeader("role_id") == config.RoleClient {
		h.handleResponse(c, http.OK, struct{}{})
		return
	}

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.PaymentService().GetList(
		context.Background(),
		&users_service.GetListPaymentRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdatePayment godoc
// @ID update_payment
// @Router /payment/{id} [PUT]
// @Summary Update Payment
// @Description Update Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdatePayment true "UpdatePayment"
// @Success 200 {object} http.Response{data=users_service.Payment} "Payment data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePayment(c *gin.Context) {

	var payment users_service.UpdatePayment

	payment.Id = c.Param("id")

	if !util.IsValidUUID(payment.Id) {
		h.handleResponse(c, http.InvalidArgument, "payment id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&payment)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PaymentService().Update(
		c.Request.Context(),
		&payment,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeletePayment godoc
// @ID delete_payment
// @Router /payment/{id} [DELETE]
// @Summary Delete Payment
// @Description Delete Payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Payment data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePayment(c *gin.Context) {

	paymentId := c.Param("id")

	if !util.IsValidUUID(paymentId) {
		h.handleResponse(c, http.InvalidArgument, "payment id is an invalid uuid")
		return
	}

	resp, err := h.services.PaymentService().Delete(
		c.Request.Context(),
		&users_service.PaymentPrimaryKey{Id: paymentId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
