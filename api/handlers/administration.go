package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateAdministration godoc
// @ID create_administration
// @Router /administration [POST]
// @Summary Create Administration
// @Description  Create Administration
// @Tags Administration
// @Accept json
// @Produce json
// @Param profile body users_service.CreateAdministration true "CreateAdministrationBody"
// @Success 200 {object} http.Response{data=users_service.Administration} "GetAdministrationBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdministration(c *gin.Context) {

	var administration users_service.CreateAdministration

	err := c.ShouldBindJSON(&administration)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AdministrationService().Create(
		c.Request.Context(),
		&administration,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAdministrationByID godoc
// @ID get_administration_by_id
// @Router /administration/{id} [GET]
// @Summary Get Administration  By ID
// @Description Get Administration  By ID
// @Tags Administration
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Administration} "AdministrationBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdministrationByID(c *gin.Context) {

	administrationID := c.Param("id")

	if !util.IsValidUUID(administrationID) {
		h.handleResponse(c, http.InvalidArgument, "administration id is an invalid uuid")
		return
	}

	resp, err := h.services.AdministrationService().GetById(
		context.Background(),
		&users_service.AdministrationPrimaryKey{
			Id: administrationID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetAdministrationList godoc
// @ID get_administration_list
// @Router /administration [GET]
// @Summary Get Administration s List
// @Description  Get Administration s List
// @Tags Administration
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param first_name query string false "full_name"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListAdministrationResponse} "GetAllAdministrationResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdministrationList(c *gin.Context) {

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

	resp, err := h.services.AdministrationService().GetList(
		context.Background(),
		&users_service.GetListAdministrationRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("full_name"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateAdministration godoc
// @ID update_administration
// @Router /administration/{id} [PUT]
// @Summary Update Administration
// @Description Update Administration
// @Tags Administration
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateAdministration true "UpdateAdministration"
// @Success 200 {object} http.Response{data=users_service.Administration} "Administration data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdministration(c *gin.Context) {

	var administration users_service.UpdateAdministration

	administration.Id = c.Param("id")

	if !util.IsValidUUID(administration.Id) {
		h.handleResponse(c, http.InvalidArgument, "administration id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&administration)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AdministrationService().Update(
		c.Request.Context(),
		&administration,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAdministration godoc
// @ID delete_administration
// @Router /administration/{id} [DELETE]
// @Summary Delete Administration
// @Description Delete Administration
// @Tags Administration
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Administration data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAdministration(c *gin.Context) {

	administrationId := c.Param("id")

	if !util.IsValidUUID(administrationId) {
		h.handleResponse(c, http.InvalidArgument, "administration id is an invalid uuid")
		return
	}

	resp, err := h.services.AdministrationService().Delete(
		c.Request.Context(),
		&users_service.AdministrationPrimaryKey{Id: administrationId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
