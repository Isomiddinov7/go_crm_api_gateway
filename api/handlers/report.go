package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// GetReportByID godoc
// @ID get_report_by_id
// @Router /report/{id} [GET]
// @Summary Get Report  By ID
// @Description Get Report  By ID
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.StudentList} "ReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudent(c *gin.Context) {

	reportID := c.Param("id")

	if !util.IsValidUUID(reportID) {
		h.handleResponse(c, http.InvalidArgument, "report id is an invalid uuid")
		return
	}

	resp, err := h.services.ReportService().GetStudent(
		context.Background(),
		&users_service.ReportRequest{
			BranchId: reportID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetReportByID godoc
// @ID get_report_by_id
// @Router /report/{id} [GET]
// @Summary Get Report  By ID
// @Description Get Report  By ID
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.SupportTeacherList} "ReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSupportTeacher(c *gin.Context) {

	reportID := c.Param("id")

	if !util.IsValidUUID(reportID) {
		h.handleResponse(c, http.InvalidArgument, "report id is an invalid uuid")
		return
	}

	resp, err := h.services.ReportService().GetSupportTeacher(
		context.Background(),
		&users_service.ReportRequest{
			BranchId: reportID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetReportByID godoc
// @ID get_report_by_id
// @Router /report/{id} [GET]
// @Summary Get Report  By ID
// @Description Get Report  By ID
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.AdminList} "ReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdministrator(c *gin.Context) {

	reportID := c.Param("id")

	if !util.IsValidUUID(reportID) {
		h.handleResponse(c, http.InvalidArgument, "report id is an invalid uuid")
		return
	}

	resp, err := h.services.ReportService().GetAdministrator(
		context.Background(),
		&users_service.ReportRequest{
			BranchId: reportID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetReportByID godoc
// @ID get_report_by_id
// @Router /report/{id} [GET]
// @Summary Get Report  By ID
// @Description Get Report  By ID
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.TeacherList} "ReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTeacher(c *gin.Context) {

	reportID := c.Param("id")

	if !util.IsValidUUID(reportID) {
		h.handleResponse(c, http.InvalidArgument, "report id is an invalid uuid")
		return
	}

	resp, err := h.services.ReportService().GetTeacher(
		context.Background(),
		&users_service.ReportRequest{
			BranchId: reportID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetReportByID godoc
// @ID get_report_by_id
// @Router /report/{id} [GET]
// @Summary Get Report  By ID
// @Description Get Report  By ID
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.StudentList} "ReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentGroup(c *gin.Context) {

	reportID := c.Param("id")

	if !util.IsValidUUID(reportID) {
		h.handleResponse(c, http.InvalidArgument, "report id is an invalid uuid")
		return
	}

	resp, err := h.services.ReportService().GetStudentGroup(
		context.Background(),
		&users_service.GroupId{
			GroupId: reportID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetReportByID godoc
// @ID get_report_by_id
// @Router /report/{id} [GET]
// @Summary Get Report  By ID
// @Description Get Report  By ID
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.StudentReportByTeacher} "ReportBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetStudentGroupByTeacher(c *gin.Context) {

	reportID := c.Param("id")

	if !util.IsValidUUID(reportID) {
		h.handleResponse(c, http.InvalidArgument, "report id is an invalid uuid")
		return
	}

	resp, err := h.services.ReportService().GetStudentGroup(
		context.Background(),
		&users_service.GroupId{
			GroupId: reportID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
