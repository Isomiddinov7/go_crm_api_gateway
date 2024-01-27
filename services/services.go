package services

import (
	"go_crm_api_gateway/config"
	"go_crm_api_gateway/genproto/users_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	TeacherService() users_service.TeacherServiceClient
	SupportTeacherService() users_service.SupportTeacherServiceClient
	AdministrationService() users_service.AdministrationServiceClient
	BranchService() users_service.BranchServiceClient
	EventService() users_service.EventServiceClient
	GroupService() users_service.GroupServiceClient
	JournalService() users_service.JournalServiceClient
	LessonService() users_service.LessonServiceClient
	ManagerService() users_service.ManagerServiceClient
	SchedulesService() users_service.SchedulesServiceClient
	StudentService() users_service.StudentServiceClient
	TaskService() users_service.TaskServiceClient
	RoleService() users_service.RoleServiceClient
	AssignStudentService() users_service.AssignStudentServiceClient
	PaymentService() users_service.PaymentServiceClient
	ReportService() users_service.ReportServiceClient
	ScoreService() users_service.ScoreServiceClient
}

type grpcClients struct {
	teacherService         users_service.TeacherServiceClient
	support_teacherService users_service.SupportTeacherServiceClient
	administrationService  users_service.AdministrationServiceClient
	branchService          users_service.BranchServiceClient
	eventService           users_service.EventServiceClient
	groupService           users_service.GroupServiceClient
	journalService         users_service.JournalServiceClient
	lessonService          users_service.LessonServiceClient
	managerService         users_service.ManagerServiceClient
	schedulesService       users_service.SchedulesServiceClient
	studentService         users_service.StudentServiceClient
	taskService            users_service.TaskServiceClient
	roleService            users_service.RoleServiceClient
	assignStudentService   users_service.AssignStudentServiceClient
	paymentService         users_service.PaymentServiceClient
	reportService          users_service.ReportServiceClient
	scoreService           users_service.ScoreServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	// User Service...
	connUsersService, err := grpc.Dial(
		cfg.UsersServiceHost+cfg.UsersGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		teacherService:         users_service.NewTeacherServiceClient(connUsersService),
		support_teacherService: users_service.NewSupportTeacherServiceClient(connUsersService),
		administrationService:  users_service.NewAdministrationServiceClient(connUsersService),
		branchService:          users_service.NewBranchServiceClient(connUsersService),
		eventService:           users_service.NewEventServiceClient(connUsersService),
		groupService:           users_service.NewGroupServiceClient(connUsersService),
		journalService:         users_service.NewJournalServiceClient(connUsersService),
		lessonService:          users_service.NewLessonServiceClient(connUsersService),
		managerService:         users_service.NewManagerServiceClient(connUsersService),
		schedulesService:       users_service.NewSchedulesServiceClient(connUsersService),
		studentService:         users_service.NewStudentServiceClient(connUsersService),
		taskService:            users_service.NewTaskServiceClient(connUsersService),
		roleService:            users_service.NewRoleServiceClient(connUsersService),
		assignStudentService:   users_service.NewAssignStudentServiceClient(connUsersService),
		paymentService:         users_service.NewPaymentServiceClient(connUsersService),
		reportService:          users_service.NewReportServiceClient(connUsersService),
		scoreService:           users_service.NewScoreServiceClient(connUsersService),
	}, nil
}

func (g *grpcClients) TeacherService() users_service.TeacherServiceClient {
	return g.teacherService
}

func (g *grpcClients) SupportTeacherService() users_service.SupportTeacherServiceClient {
	return g.support_teacherService
}

func (g *grpcClients) AdministrationService() users_service.AdministrationServiceClient {
	return g.administrationService
}

func (g *grpcClients) BranchService() users_service.BranchServiceClient {
	return g.branchService
}

func (g *grpcClients) EventService() users_service.EventServiceClient {
	return g.eventService
}

func (g *grpcClients) GroupService() users_service.GroupServiceClient {
	return g.groupService
}

func (g *grpcClients) JournalService() users_service.JournalServiceClient {
	return g.journalService
}

func (g *grpcClients) LessonService() users_service.LessonServiceClient {
	return g.lessonService
}

func (g *grpcClients) ManagerService() users_service.ManagerServiceClient {
	return g.managerService
}

func (g *grpcClients) SchedulesService() users_service.SchedulesServiceClient {
	return g.schedulesService
}

func (g *grpcClients) StudentService() users_service.StudentServiceClient {
	return g.studentService
}

func (g *grpcClients) TaskService() users_service.TaskServiceClient {
	return g.taskService
}

func (g *grpcClients) RoleService() users_service.RoleServiceClient {
	return g.roleService
}

func (g *grpcClients) AssignStudentService() users_service.AssignStudentServiceClient {
	return g.assignStudentService
}

func (g *grpcClients) PaymentService() users_service.PaymentServiceClient {
	return g.paymentService
}

func (g *grpcClients) ReportService() users_service.ReportServiceClient {
	return g.reportService
}

func (g *grpcClients) ScoreService() users_service.ScoreServiceClient {
	return g.scoreService
}
