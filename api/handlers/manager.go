package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateManager godoc
// @ID create_manager
// @Router /manager [POST]
// @Summary Create Manager
// @Description  Create Manager
// @Tags Manager
// @Accept json
// @Produce json
// @Param profile body users_service.CreateManager true "CreateManagerBody"
// @Success 200 {object} http.Response{data=users_service.Manager} "GetManagerBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateManager(c *gin.Context) {

	var manager users_service.CreateManager

	err := c.ShouldBindJSON(&manager)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ManagerService().Create(
		c.Request.Context(),
		&manager,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetManagerByID godoc
// @ID get_manager_by_id
// @Router /manager/{id} [GET]
// @Summary Get Manager  By ID
// @Description Get Manager  By ID
// @Tags Manager
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Manager} "ManagerBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetManagerByID(c *gin.Context) {

	managerID := c.Param("id")

	if !util.IsValidUUID(managerID) {
		h.handleResponse(c, http.InvalidArgument, "manager id is an invalid uuid")
		return
	}

	resp, err := h.services.ManagerService().GetById(
		context.Background(),
		&users_service.ManagerPrimaryKey{
			Id: managerID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetManagerList godoc
// @ID get_manager_list
// @Router /manager [GET]
// @Summary Get Manager s List
// @Description  Get Manager s List
// @Tags Manager
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param first_name query string false "full_name"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListManagerResponse} "GetAllManagerResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetManagerList(c *gin.Context) {

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

	resp, err := h.services.ManagerService().GetList(
		context.Background(),
		&users_service.GetListManagerRequest{
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

// UpdateManager godoc
// @ID update_manager
// @Router /manager/{id} [PUT]
// @Summary Update Manager
// @Description Update Manager
// @Tags Manager
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateManager true "UpdateManager"
// @Success 200 {object} http.Response{data=users_service.Manager} "Manager data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateManager(c *gin.Context) {

	var manager users_service.UpdateManager

	manager.Id = c.Param("id")

	if !util.IsValidUUID(manager.Id) {
		h.handleResponse(c, http.InvalidArgument, "manager id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&manager)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ManagerService().Update(
		c.Request.Context(),
		&manager,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteManager godoc
// @ID delete_manager
// @Router /manager/{id} [DELETE]
// @Summary Delete Manager
// @Description Delete Manager
// @Tags Manager
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Manager data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteManager(c *gin.Context) {

	managerId := c.Param("id")

	if !util.IsValidUUID(managerId) {
		h.handleResponse(c, http.InvalidArgument, "manager id is an invalid uuid")
		return
	}

	resp, err := h.services.ManagerService().Delete(
		c.Request.Context(),
		&users_service.ManagerPrimaryKey{Id: managerId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
