package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateGroup godoc
// @ID create_group
// @Router /group [POST]
// @Summary Create Group
// @Description  Create Group
// @Tags Group
// @Accept json
// @Produce json
// @Param profile body users_service.CreateGroup true "CreateGroupBody"
// @Success 200 {object} http.Response{data=users_service.Group} "GetGroupBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateGroup(c *gin.Context) {

	var group users_service.CreateGroup

	err := c.ShouldBindJSON(&group)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.GroupService().Create(
		c.Request.Context(),
		&group,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetGroupByID godoc
// @ID get_group_by_id
// @Router /group/{id} [GET]
// @Summary Get Group  By ID
// @Description Get Group  By ID
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Group} "GroupBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGroupByID(c *gin.Context) {

	groupID := c.Param("id")

	if !util.IsValidUUID(groupID) {
		h.handleResponse(c, http.InvalidArgument, "group id is an invalid uuid")
		return
	}

	resp, err := h.services.GroupService().GetById(
		context.Background(),
		&users_service.GroupPrimaryKey{
			Id: groupID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetGroupList godoc
// @ID get_group_list
// @Router /group [GET]
// @Summary Get Group s List
// @Description  Get Group s List
// @Tags Group
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param first_name query string false "full_name"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListGroupResponse} "GetAllGroupResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGroupList(c *gin.Context) {

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

	resp, err := h.services.GroupService().GetList(
		context.Background(),
		&users_service.GetListGroupRequest{
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

// UpdateGroup godoc
// @ID update_group
// @Router /group/{id} [PUT]
// @Summary Update Group
// @Description Update Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateGroup true "UpdateGroup"
// @Success 200 {object} http.Response{data=users_service.Group} "Group data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateGroup(c *gin.Context) {

	var group users_service.UpdateGroup

	group.Id = c.Param("id")

	if !util.IsValidUUID(group.Id) {
		h.handleResponse(c, http.InvalidArgument, "group id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&group)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.GroupService().Update(
		c.Request.Context(),
		&group,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteGroup godoc
// @ID delete_group
// @Router /group/{id} [DELETE]
// @Summary Delete Group
// @Description Delete Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Group data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteGroup(c *gin.Context) {

	groupId := c.Param("id")

	if !util.IsValidUUID(groupId) {
		h.handleResponse(c, http.InvalidArgument, "group id is an invalid uuid")
		return
	}

	resp, err := h.services.GroupService().Delete(
		c.Request.Context(),
		&users_service.GroupPrimaryKey{Id: groupId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
