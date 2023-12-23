package app

import "go-do/internal/services"

var (
	healthService = services.NewHealthService(&s)
	pingService   = services.NewPingService(&s)
	taskService   = services.NewTaskService(&s)
	userService   = services.NewUserService(&s)
)

var service = []services.Service{
	&healthService,
	&pingService,
	&taskService,
	&userService,
}
