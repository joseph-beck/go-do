package app

import (
	"go-do/internal/services"

	routey "github.com/joseph-beck/routey/pkg/router"
)

var (
	healthService = services.NewHealthService(&s)
	pingService   = services.NewPingService(&s)
	taskService   = services.NewTaskService(&s)
	userService   = services.NewUserService(&s)
)

var service = []routey.Service{
	&healthService,
	&pingService,
	&taskService,
	&userService,
}
