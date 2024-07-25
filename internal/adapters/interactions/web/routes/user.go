package routes

import (
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web/middleware"
	"github.com/Peterwmoss/LiCa/internal/core/domain/ports"
)

type UserRouter struct {
  Auth middleware.Middleware
  Service ports.UserService
}

func NewUserRouter(auth middleware.Middleware, service ports.UserService) UserRouter {
  return UserRouter{
    Auth: auth,
    Service: service,
  }
}
