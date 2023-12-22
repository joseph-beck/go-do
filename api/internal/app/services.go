package app

import "go-do/internal/services"

var (
	healthService = services.NewHealthService()
	pingService   = services.NewPingService()
	taskService   = services.NewTaskService()
	userService   = services.NewUserService()
)

var service = []services.Service{
	&healthService,
	&pingService,
	&taskService,
	&userService,
}
