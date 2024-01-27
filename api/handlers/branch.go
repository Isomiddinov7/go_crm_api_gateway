package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @ID create_branch
// @Router /branch [POST]
// @Summary Create Branch
// @Description  Create Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param profile body users_service.CreateBranch true "CreateBranchBody"
// @Success 200 {object} http.Response{data=users_service.Branch} "GetBranchBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBranch(c *gin.Context) {

	var branch users_service.CreateBranch

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchService().Create(
		c.Request.Context(),
		&branch,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetBranchByID godoc
// @ID get_branch_by_id
// @Router /branch/{id} [GET]
// @Summary Get Branch  By ID
// @Description Get Branch  By ID
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Branch} "BranchBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBranchByID(c *gin.Context) {

	branchID := c.Param("id")

	if !util.IsValidUUID(branchID) {
		h.handleResponse(c, http.InvalidArgument, "branch id is an invalid uuid")
		return
	}

	resp, err := h.services.BranchService().GetById(
		context.Background(),
		&users_service.BranchPrimaryKey{
			Id: branchID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetBranchList godoc
// @ID get_branch_list
// @Router /branch [GET]
// @Summary Get Branch s List
// @Description  Get Branch s List
// @Tags Branch
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListBranchResponse} "GetAllBranchResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBranchList(c *gin.Context) {

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

	resp, err := h.services.BranchService().GetList(
		context.Background(),
		&users_service.GetListBranchRequest{
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

// UpdateBranch godoc
// @ID update_branch
// @Router /branch/{id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateBranch true "UpdateBranch"
// @Success 200 {object} http.Response{data=users_service.Branch} "Branch data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBranch(c *gin.Context) {

	var branch users_service.UpdateBranch

	branch.Id = c.Param("id")

	if !util.IsValidUUID(branch.Id) {
		h.handleResponse(c, http.InvalidArgument, "branch id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchService().Update(
		c.Request.Context(),
		&branch,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteBranch godoc
// @ID delete_branch
// @Router /branch/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Branch data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBranch(c *gin.Context) {

	branchId := c.Param("id")

	if !util.IsValidUUID(branchId) {
		h.handleResponse(c, http.InvalidArgument, "branch id is an invalid uuid")
		return
	}

	resp, err := h.services.BranchService().Delete(
		c.Request.Context(),
		&users_service.BranchPrimaryKey{Id: branchId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
