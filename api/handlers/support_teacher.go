package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateSupportTeacher godoc
// @ID create_support_teacher
// @Router /support_teacher [POST]
// @Summary Create SupportTeacher
// @Description  Create SupportTeacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param profile body users_service.CreateSupportTeacher true "CreateSupportTeacherBody"
// @Success 200 {object} http.Response{data=users_service.SupportTeacher} "GetSupportTeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSupportTeacher(c *gin.Context) {

	var support_teacher users_service.CreateSupportTeacher

	err := c.ShouldBindJSON(&support_teacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SupportTeacherService().Create(
		c.Request.Context(),
		&support_teacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSupportTeacherByID godoc
// @ID get_support_teacher_by_id
// @Router /support_teacher/{id} [GET]
// @Summary Get SupportTeacher  By ID
// @Description Get SupportTeacher  By ID
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.SupportTeacher} "SupportTeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSupportTeacherByID(c *gin.Context) {

	support_teacherID := c.Param("id")

	if !util.IsValidUUID(support_teacherID) {
		h.handleResponse(c, http.InvalidArgument, "support_teacher id is an invalid uuid")
		return
	}

	resp, err := h.services.SupportTeacherService().GetById(
		context.Background(),
		&users_service.SupportTeacherPrimaryKey{
			Id: support_teacherID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetSupportTeacherList godoc
// @ID get_support_teacher_list
// @Router /support_teacher [GET]
// @Summary Get SupportTeacher s List
// @Description  Get SupportTeacher s List
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param first_name query string false "full_name"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListSupportTeacherResponse} "GetAllSupportTeacherResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSupportTeacherList(c *gin.Context) {

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

	resp, err := h.services.SupportTeacherService().GetList(
		context.Background(),
		&users_service.GetListSupportTeacherRequest{
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

// UpdateSupportTeacher godoc
// @ID update_support_teacher
// @Router /support_teacher/{id} [PUT]
// @Summary Update SupportTeacher
// @Description Update SupportTeacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateSupportTeacher true "UpdateSupportTeacher"
// @Success 200 {object} http.Response{data=users_service.SupportTeacher} "SupportTeacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSupportTeacher(c *gin.Context) {

	var support_teacher users_service.UpdateSupportTeacher

	support_teacher.Id = c.Param("id")

	if !util.IsValidUUID(support_teacher.Id) {
		h.handleResponse(c, http.InvalidArgument, "support_teacher id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&support_teacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SupportTeacherService().Update(
		c.Request.Context(),
		&support_teacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteSupportTeacher godoc
// @ID delete_support_teacher
// @Router /support_teacher/{id} [DELETE]
// @Summary Delete SupportTeacher
// @Description Delete SupportTeacher
// @Tags SupportTeacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "SupportTeacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSupportTeacher(c *gin.Context) {

	support_teacherId := c.Param("id")

	if !util.IsValidUUID(support_teacherId) {
		h.handleResponse(c, http.InvalidArgument, "support_teacher id is an invalid uuid")
		return
	}

	resp, err := h.services.SupportTeacherService().Delete(
		c.Request.Context(),
		&users_service.SupportTeacherPrimaryKey{Id: support_teacherId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
