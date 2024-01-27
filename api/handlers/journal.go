package handlers

import (
	"context"
	"go_crm_api_gateway/api/http"
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"
	"go_crm_api_gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateJournal godoc
// @ID create_journal
// @Router /journal [POST]
// @Summary Create Journal
// @Description  Create Journal
// @Tags Journal
// @Accept json
// @Produce json
// @Param profile body users_service.CreateJournal true "CreateJournalBody"
// @Success 200 {object} http.Response{data=users_service.Journal} "GetJournalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJournal(c *gin.Context) {

	var journal users_service.CreateJournal

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.JournalService().Create(
		c.Request.Context(),
		&journal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetJournalByID godoc
// @ID get_journal_by_id
// @Router /journal/{id} [GET]
// @Summary Get Journal  By ID
// @Description Get Journal  By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=users_service.Journal} "JournalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalByID(c *gin.Context) {

	journalID := c.Param("id")

	if !util.IsValidUUID(journalID) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	resp, err := h.services.JournalService().GetById(
		context.Background(),
		&users_service.JournalPrimaryKey{
			Id: journalID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Security ApiKeyAuth
// GetJournalList godoc
// @ID get_journal_list
// @Router /journal [GET]
// @Summary Get Journal s List
// @Description  Get Journal s List
// @Tags Journal
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param Platform-Id header string true "Platform-Id" default(a1924766-a9ee-11ed-afa1-0242ac120001)
// @Success 200 {object} http.Response{data=users_service.GetListJournalResponse} "GetAllJournalResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalList(c *gin.Context) {

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

	resp, err := h.services.JournalService().GetList(
		context.Background(),
		&users_service.GetListJournalRequest{
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

// UpdateJournal godoc
// @ID update_journal
// @Router /journal/{id} [PUT]
// @Summary Update Journal
// @Description Update Journal
// @Tags Journal
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body users_service.UpdateJournal true "UpdateJournal"
// @Success 200 {object} http.Response{data=users_service.Journal} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournal(c *gin.Context) {

	var journal users_service.UpdateJournal

	journal.Id = c.Param("id")

	if !util.IsValidUUID(journal.Id) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.JournalService().Update(
		c.Request.Context(),
		&journal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteJournal godoc
// @ID delete_journal
// @Router /journal/{id} [DELETE]
// @Summary Delete Journal
// @Description Delete Journal
// @Tags Journal
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournal(c *gin.Context) {

	journalId := c.Param("id")

	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	resp, err := h.services.JournalService().Delete(
		c.Request.Context(),
		&users_service.JournalPrimaryKey{Id: journalId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
