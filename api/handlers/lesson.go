package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateLesson godoc
// @ID create_lesson
// @Router /lesson [POST]
// @Summary Create Lesson
// @Description  Create Lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param profile body users_service.CreateLesson true "CreateLessonBody"
// @Success 200 {object} http.Response{data=users_service.Lesson} "GetLessonBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateLesson(c *gin.Context) {

	var lesson users_service.CreateLesson

	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.LessonService().Create(
		c.Request.Context(),
		&lesson,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetLessonByID godoc
// @ID get_lesson_by_id
// @Router /lesson/{id} [GET]
// @Summary Get Lesson  By ID
// @Description Get Lesson  By ID
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Lesson} "LessonBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetLessonByID(c *gin.Context) {

	lessonID := c.Param("id")

	if !util.IsValidUUID(lessonID) {
		h.handleResponse(c, http.InvalidArgument, "lesson id is an invalid uuid")
		return
	}

	resp, err := h.services.LessonService().GetById(
		context.Background(),
		&users_service.LessonPrimaryKey{
			Id: lessonID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetLessonList godoc
// @ID get_lesson_list
// @Router /lesson [GET]
// @Summary Get Lesson s List
// @Description  Get Lesson s List
// @Tags Lesson
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListLessonResponse} "GetAllLessonResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetLessonList(c *gin.Context) {

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

	resp, err := h.services.LessonService().GetList(
		context.Background(),
		&users_service.GetListLessonRequest{
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

// UpdateLesson godoc
// @ID update_lesson
// @Router /lesson/{id} [PUT]
// @Summary Update Lesson
// @Description Update Lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateLesson true "UpdateLesson"
// @Success 200 {object} http.Response{data=users_service.Lesson} "Lesson data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateLesson(c *gin.Context) {

	var lesson users_service.UpdateLesson

	lesson.Id = c.Param("id")

	if !util.IsValidUUID(lesson.Id) {
		h.handleResponse(c, http.InvalidArgument, "lesson id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.LessonService().Update(
		c.Request.Context(),
		&lesson,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteLesson godoc
// @ID delete_lesson
// @Router /lesson/{id} [DELETE]
// @Summary Delete Lesson
// @Description Delete Lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Lesson data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteLesson(c *gin.Context) {

	lessonId := c.Param("id")

	if !util.IsValidUUID(lessonId) {
		h.handleResponse(c, http.InvalidArgument, "lesson id is an invalid uuid")
		return
	}

	resp, err := h.services.LessonService().Delete(
		c.Request.Context(),
		&users_service.LessonPrimaryKey{Id: lessonId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
