package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @ID create_student
// @Router /student [POST]
// @Summary Create Student
// @Description  Create Student
// @Tags Student
// @Accept json
// @Produce json
// @Param profile body users_service.CreateStudent true "CreateStudentBody"
// @Success 200 {object} http.Response{data=users_service.Student} "GetStudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateStudent(c *gin.Context) {

	var student users_service.CreateStudent

	err := c.ShouldBindJSON(&student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.StudentService().Create(
		c.Request.Context(),
		&student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetStudentByID godoc
// @ID get_student_by_id
// @Router /student/{id} [GET]
// @Summary Get Student  By ID
// @Description Get Student  By ID
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Student} "StudentBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentByID(c *gin.Context) {

	studentID := c.Param("id")

	if !util.IsValidUUID(studentID) {
		h.handleResponse(c, http.InvalidArgument, "student id is an invalid uuid")
		return
	}

	resp, err := h.services.StudentService().GetById(
		context.Background(),
		&users_service.StudentPrimaryKey{
			Id: studentID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetStudentList godoc
// @ID get_student_list
// @Router /student [GET]
// @Summary Get Student s List
// @Description  Get Student s List
// @Tags Student
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param first_name query string false "full_name"
// @Param phone query string false "phone"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListStudentResponse} "GetAllStudentResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentList(c *gin.Context) {

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

	resp, err := h.services.StudentService().GetList(
		context.Background(),
		&users_service.GetListStudentRequest{
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

// UpdateStudent godoc
// @ID update_student
// @Router /student/{id} [PUT]
// @Summary Update Student
// @Description Update Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateStudent true "UpdateStudent"
// @Success 200 {object} http.Response{data=users_service.Student} "Student data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateStudent(c *gin.Context) {

	var student users_service.UpdateStudent

	student.Id = c.Param("id")

	if !util.IsValidUUID(student.Id) {
		h.handleResponse(c, http.InvalidArgument, "student id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&student)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.StudentService().Update(
		c.Request.Context(),
		&student,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteStudent godoc
// @ID delete_student
// @Router /student/{id} [DELETE]
// @Summary Delete Student
// @Description Delete Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Student data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteStudent(c *gin.Context) {

	studentId := c.Param("id")

	if !util.IsValidUUID(studentId) {
		h.handleResponse(c, http.InvalidArgument, "student id is an invalid uuid")
		return
	}

	resp, err := h.services.StudentService().Delete(
		c.Request.Context(),
		&users_service.StudentPrimaryKey{Id: studentId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
