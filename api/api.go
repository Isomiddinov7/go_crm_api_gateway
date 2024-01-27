package api

import (
	"fmt"
	"go_crm_api_gateway/api/docs"
	"go_crm_api_gateway/api/handlers"
	"go_crm_api_gateway/config"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
const (
	RoleTeacher = "teacher"
	RoleStudent = "student"
	RoleAdmin   = "admin"
	// Add other roles as needed
)

func TeacherOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated and has the role of a teacher
		if userRole := c.GetHeader("Role"); userRole != RoleTeacher {
			fmt.Println()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customCORSMiddleware())
	r.Use(MaxAllowed(5000))
	// r.Use(h.CheckPasswordMiddleware())

	// USERS

	r.POST("/teacher", h.CreateTeacher)
	r.GET("/teacher/:id", h.GetTeacherByID)
	r.GET("/teacher", h.GetTeacherList)
	r.PUT("/teacher/:id", h.UpdateTeacher)
	r.DELETE("/teacher/:id", h.DeleteTeacher)

	r.POST("/support_teacher", h.CreateSupportTeacher)
	r.GET("/support_teacher/:id", h.GetSupportTeacherByID)
	r.GET("/support_teacher", h.GetSupportTeacherList)
	r.PUT("/support_teacher/:id", h.UpdateSupportTeacher)
	r.DELETE("/support_teacher/:id", h.DeleteSupportTeacher)

	r.POST("/administration", h.CreateAdministration)
	r.GET("/administration/:id", h.GetAdministrationByID)
	r.GET("/administration", h.GetAdministrationList)
	r.PUT("/administration/:id", h.UpdateAdministration)
	r.DELETE("/administration/:id", h.DeleteAdministration)

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranchByID)
	r.GET("/branch", h.GetBranchList)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/event", h.CreateEvent)
	r.GET("/event/:id", h.GetEventByID)
	r.GET("/event", h.GetEventList)
	r.PUT("/event/:id", h.UpdateEvent)
	r.DELETE("/event/:id", h.DeleteEvent)

	r.POST("/group", h.CreateGroup)
	r.GET("/group/:id", h.GetGroupByID)
	r.GET("/group", h.GetGroupList)
	r.PUT("/group/:id", h.UpdateGroup)
	r.DELETE("/group/:id", h.DeleteGroup)

	r.POST("/journal", h.CreateJournal)
	r.GET("/journal/:id", h.GetJournalByID)
	r.GET("/journal", h.GetJournalList)
	r.PUT("/journal/:id", h.UpdateJournal)
	r.DELETE("/journal/:id", h.DeleteJournal)

	r.POST("/lesson", h.CreateLesson)
	r.GET("/lesson/:id", h.GetLessonByID)
	r.GET("/lesson", h.GetLessonList)
	r.PUT("/lesson/:id", h.UpdateLesson)
	r.DELETE("/lesson/:id", h.DeleteLesson)

	r.POST("/manager", h.CreateManager)
	r.GET("/manager/:id", h.GetManagerByID)
	r.GET("/manager", h.GetManagerList)
	r.PUT("/manager/:id", h.UpdateManager)
	r.DELETE("/manager/:id", h.DeleteManager)

	r.POST("/student", h.CreateStudent)
	r.GET("/student/:id", h.GetStudentByID)
	r.GET("/student", h.GetStudentList)
	r.PUT("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)

	r.POST("/task", h.CreateTask)
	r.GET("/task/:id", h.GetTaskByID)
	r.GET("/task", h.GetTaskList)
	r.PUT("/task/:id", h.UpdateTask)
	r.DELETE("/task/:id", h.DeleteTask)

	r.POST("/assign_student", h.CreateAssignStudent)
	r.GET("/assign_student/:id", h.GetAssignStudentByID)
	r.GET("/assign_student", h.GetAssignStudentList)
	r.PUT("/assign_student/:id", h.UpdateAssignStudent)
	r.DELETE("/assign_student/:id", h.DeleteAssignStudent)

	r.POST("/payment", h.CreatePayment)
	r.GET("/payment/:id", h.GetPaymentByID)
	r.GET("/payment", h.GetPaymentList)
	r.PUT("/payment/:id", h.UpdatePayment)
	r.DELETE("/payment/:id", h.DeletePayment)

	r.POST("/score", h.CreateScore)
	r.GET("/score/:id", h.GetScoreByID)
	r.GET("/score", h.GetScoreList)
	r.PUT("/score/:id", h.UpdateScore)
	r.DELETE("/score/:id", h.DeleteScore)

	//report
	r.GET("/report/get_student/:id", h.GetStudent)
	r.GET("/report/get_support_teacher/:id", h.GetSupportTeacher)
	r.GET("/report/get_administrator/:id", h.GetAdministrator)
	r.GET("/report/get_teacher/:id", h.GetTeacher)
	r.GET("/report/get_student_group/:id", h.GetStudentGroup)
	r.GET("/report/get_student_group_by_teacher/:id", h.GetStudentGroupByTeacher)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}
