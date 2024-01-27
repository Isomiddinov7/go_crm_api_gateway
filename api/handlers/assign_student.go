package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateAssignStudent godoc
// @ID create_assign_student
// @Router /assign_student [POST]
// @Summary Create AssignStudent
// @Description  Create AssignStudent
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param profile body users_service.CreateAssignStudent true "CreateAssignStudentBody"
// @Success 200 {object} http.Response{data=users_service.AssignStudent} "GetAssignStudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAssignStudent(c *gin.Context) {

	var assign_student users_service.CreateAssignStudent

	err := c.ShouldBindJSON(&assign_student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AssignStudentService().Create(
		c.Request.Context(),
		&assign_student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAssignStudentByID godoc
// @ID get_assign_student_by_id
// @Router /assign_student/{id} [GET]
// @Summary Get AssignStudent  By ID
// @Description Get AssignStudent  By ID
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.AssignStudent} "AssignStudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAssignStudentByID(c *gin.Context) {

	assign_studentID := c.Param("id")

	if !util.IsValidUUID(assign_studentID) {
		h.handleResponse(c, http.InvalidArgument, "assign_student id is an invalid uuid")
		return
	}

	resp, err := h.services.AssignStudentService().GetById(
		context.Background(),
		&users_service.AssignStudentPrimaryKey{
			Id: assign_studentID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetAssignStudentList godoc
// @ID get_assign_student_list
// @Router /assign_student [GET]
// @Summary Get AssignStudent s List
// @Description  Get AssignStudent s List
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListAssignStudentResponse} "GetAllAssignStudentResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAssignStudentList(c *gin.Context) {

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

	resp, err := h.services.AssignStudentService().GetList(
		context.Background(),
		&users_service.GetListAssignStudentRequest{
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

// UpdateAssignStudent godoc
// @ID update_assign_student
// @Router /assign_student/{id} [PUT]
// @Summary Update AssignStudent
// @Description Update AssignStudent
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateAssignStudent true "UpdateAssignStudent"
// @Success 200 {object} http.Response{data=users_service.AssignStudent} "AssignStudent data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAssignStudent(c *gin.Context) {

	var assign_student users_service.UpdateAssignStudent

	assign_student.Id = c.Param("id")

	if !util.IsValidUUID(assign_student.Id) {
		h.handleResponse(c, http.InvalidArgument, "assign_student id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&assign_student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AssignStudentService().Update(
		c.Request.Context(),
		&assign_student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAssignStudent godoc
// @ID delete_assign_student
// @Router /assign_student/{id} [DELETE]
// @Summary Delete AssignStudent
// @Description Delete AssignStudent
// @Tags AssignStudent
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "AssignStudent data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAssignStudent(c *gin.Context) {

	assign_studentId := c.Param("id")

	if !util.IsValidUUID(assign_studentId) {
		h.handleResponse(c, http.InvalidArgument, "assign_student id is an invalid uuid")
		return
	}

	resp, err := h.services.AssignStudentService().Delete(
		c.Request.Context(),
		&users_service.AssignStudentPrimaryKey{Id: assign_studentId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
