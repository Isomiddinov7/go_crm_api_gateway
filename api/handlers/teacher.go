package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateTeacher godoc
// @ID create_teacher
// @Router /teacher [POST]
// @Summary Create Teacher
// @Description  Create Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param profile body users_service.CreateTeacher true "CreateTeacherBody"
// @Success 200 {object} http.Response{data=users_service.Teacher} "GetTeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateTeacher(c *gin.Context) {

	var teacher users_service.CreateTeacher

	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.TeacherService().Create(
		c.Request.Context(),
		&teacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetTeacherByID godoc
// @ID get_teacher_by_id
// @Router /teacher/{id} [GET]
// @Summary Get Teacher  By ID
// @Description Get Teacher  By ID
// @Tags Teacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Teacher} "TeacherBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTeacherByID(c *gin.Context) {

	teacherID := c.Param("id")

	if !util.IsValidUUID(teacherID) {
		h.handleResponse(c, http.InvalidArgument, "teacher id is an invalid uuid")
		return
	}

	resp, err := h.services.TeacherService().GetById(
		context.Background(),
		&users_service.TeacherPrimaryKey{
			Id: teacherID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetTeacherList godoc
// @ID get_teacher_list
// @Router /teacher [GET]
// @Summary Get Teacher s List
// @Description  Get Teacher s List
// @Tags Teacher
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param first_name query string false "full_name"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListTeacherResponse} "GetAllTeacherResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTeacherList(c *gin.Context) {

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

	resp, err := h.services.TeacherService().GetList(
		context.Background(),
		&users_service.GetListTeacherRequest{
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

// UpdateTeacher godoc
// @ID update_teacher
// @Router /teacher/{id} [PUT]
// @Summary Update Teacher
// @Description Update Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateTeacher true "UpdateTeacher"
// @Success 200 {object} http.Response{data=users_service.Teacher} "Teacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateTeacher(c *gin.Context) {

	var teacher users_service.UpdateTeacher

	teacher.Id = c.Param("id")

	if !util.IsValidUUID(teacher.Id) {
		h.handleResponse(c, http.InvalidArgument, "teacher id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.TeacherService().Update(
		c.Request.Context(),
		&teacher,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteTeacher godoc
// @ID delete_teacher
// @Router /teacher/{id} [DELETE]
// @Summary Delete Teacher
// @Description Delete Teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Teacher data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteTeacher(c *gin.Context) {

	teacherId := c.Param("id")

	if !util.IsValidUUID(teacherId) {
		h.handleResponse(c, http.InvalidArgument, "teacher id is an invalid uuid")
		return
	}

	resp, err := h.services.TeacherService().Delete(
		c.Request.Context(),
		&users_service.TeacherPrimaryKey{Id: teacherId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
