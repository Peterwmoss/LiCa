package main

import (
	"log/slog"
	"os"

	_ "github.com/Peterwmoss/LiCa/internal/env"
	_ "github.com/Peterwmoss/LiCa/internal/logger"

	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql"
	"github.com/Peterwmoss/LiCa/internal/adapters/infrastructure/postgresql/repositories"
	"github.com/Peterwmoss/LiCa/internal/adapters/interactions/web"
	"github.com/Peterwmoss/LiCa/internal/core/domain/services"
)


func main() {
	db := postgresql.Get()
	defer db.Close()

  // listRepo := repositories.NewListRepository(db)
  // listService := services.NewListService(listRepo, nil)

  userRepo := repositories.NewUserRepository(db)
  userService := services.NewUserService(userRepo)

  router := web.Router{
    UserService: userService,
  }

  if err := web.Serve(router, web.WithPort(3000)); err != nil {
		slog.Error("Failed to start api")
    os.Exit(1)
  }
}
