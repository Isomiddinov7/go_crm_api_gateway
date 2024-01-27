package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateSchedules godoc
// @ID create_schedules
// @Router /schedules [POST]
// @Summary Create Schedules
// @Description  Create Schedules
// @Tags Schedules
// @Accept json
// @Produce json
// @Param profile body users_service.CreateSchedules true "CreateSchedulesBody"
// @Success 200 {object} http.Response{data=users_service.Schedules} "GetSchedulesBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSchedules(c *gin.Context) {

	var schedules users_service.CreateSchedules

	err := c.ShouldBindJSON(&schedules)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SchedulesService().Create(
		c.Request.Context(),
		&schedules,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSchedulesByID godoc
// @ID get_schedules_by_id
// @Router /schedules/{id} [GET]
// @Summary Get Schedules  By ID
// @Description Get Schedules  By ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Schedules} "SchedulesBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSchedulesByID(c *gin.Context) {

	schedulesID := c.Param("id")

	if !util.IsValidUUID(schedulesID) {
		h.handleResponse(c, http.InvalidArgument, "schedules id is an invalid uuid")
		return
	}

	resp, err := h.services.SchedulesService().GetById(
		context.Background(),
		&users_service.SchedulesPrimaryKey{
			Id: schedulesID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetSchedulesList godoc
// @ID get_schedules_list
// @Router /schedules [GET]
// @Summary Get Schedules s List
// @Description  Get Schedules s List
// @Tags Schedules
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListSchedulesResponse} "GetAllSchedulesResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSchedulesList(c *gin.Context) {

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

	resp, err := h.services.SchedulesService().GetList(
		context.Background(),
		&users_service.GetListSchedulesRequest{
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

// UpdateSchedules godoc
// @ID update_schedules
// @Router /schedules/{id} [PUT]
// @Summary Update Schedules
// @Description Update Schedules
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateSchedules true "UpdateSchedules"
// @Success 200 {object} http.Response{data=users_service.Schedules} "Schedules data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSchedules(c *gin.Context) {

	var schedules users_service.UpdateSchedules

	schedules.Id = c.Param("id")

	if !util.IsValidUUID(schedules.Id) {
		h.handleResponse(c, http.InvalidArgument, "schedules id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&schedules)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SchedulesService().Update(
		c.Request.Context(),
		&schedules,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteSchedules godoc
// @ID delete_schedules
// @Router /schedules/{id} [DELETE]
// @Summary Delete Schedules
// @Description Delete Schedules
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Schedules data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSchedules(c *gin.Context) {

	schedulesId := c.Param("id")

	if !util.IsValidUUID(schedulesId) {
		h.handleResponse(c, http.InvalidArgument, "schedules id is an invalid uuid")
		return
	}

	resp, err := h.services.SchedulesService().Delete(
		c.Request.Context(),
		&users_service.SchedulesPrimaryKey{Id: schedulesId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
