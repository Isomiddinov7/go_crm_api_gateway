package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @ID create_event
// @Router /event [POST]
// @Summary Create Event
// @Description  Create Event
// @Tags Event
// @Accept json
// @Produce json
// @Param profile body users_service.CreateEvent true "CreateEventBody"
// @Success 200 {object} http.Response{data=users_service.Event} "GetEventBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateEvent(c *gin.Context) {

	var event users_service.CreateEvent

	err := c.ShouldBindJSON(&event)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EventService().Create(
		c.Request.Context(),
		&event,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetEventByID godoc
// @ID get_event_by_id
// @Router /event/{id} [GET]
// @Summary Get Event  By ID
// @Description Get Event  By ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Event} "EventBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEventByID(c *gin.Context) {

	eventID := c.Param("id")

	if !util.IsValidUUID(eventID) {
		h.handleResponse(c, http.InvalidArgument, "event id is an invalid uuid")
		return
	}

	resp, err := h.services.EventService().GetById(
		context.Background(),
		&users_service.EventPrimaryKey{
			Id: eventID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetEventList godoc
// @ID get_event_list
// @Router /event [GET]
// @Summary Get Event s List
// @Description  Get Event s List
// @Tags Event
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListEventResponse} "GetAllEventResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEventList(c *gin.Context) {

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

	resp, err := h.services.EventService().GetList(
		context.Background(),
		&users_service.GetListEventRequest{
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

// UpdateEvent godoc
// @ID update_event
// @Router /event/{id} [PUT]
// @Summary Update Event
// @Description Update Event
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateEvent true "UpdateEvent"
// @Success 200 {object} http.Response{data=users_service.Event} "Event data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateEvent(c *gin.Context) {

	var event users_service.UpdateEvent

	event.Id = c.Param("id")

	if !util.IsValidUUID(event.Id) {
		h.handleResponse(c, http.InvalidArgument, "event id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&event)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EventService().Update(
		c.Request.Context(),
		&event,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteEvent godoc
// @ID delete_event
// @Router /event/{id} [DELETE]
// @Summary Delete Event
// @Description Delete Event
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Event data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteEvent(c *gin.Context) {

	eventId := c.Param("id")

	if !util.IsValidUUID(eventId) {
		h.handleResponse(c, http.InvalidArgument, "event id is an invalid uuid")
		return
	}

	resp, err := h.services.EventService().Delete(
		c.Request.Context(),
		&users_service.EventPrimaryKey{Id: eventId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
