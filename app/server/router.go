package server

import (
	_ "fmt"

	"github.com/vkhoa145/go-training/app/middlewares"
	handlerUser "github.com/vkhoa145/go-training/app/modules/users/handlers"
	repositoryUser "github.com/vkhoa145/go-training/app/modules/users/repositories"
	userUseCase "github.com/vkhoa145/go-training/app/modules/users/usecase"
)

func SetupRoutes(server *Server) {
	// Auth
	userRepo := repositoryUser.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := handlerUser.NewUserHandlers(userUseCase, userRepo)

	authMiddleware := middlewares.NewAuthMiddleware(server.Config.SIGNED_STRING)

	api := server.Fiber.Group("/api/v1")
	user := api.Group("/users")
	user.Post("/signup", userHandler.SignUpUser(server.Config))
	user.Post("/signin", userHandler.SignInUser(server.Config))
	user.Get("/profile", authMiddleware, userHandler.GetUser())
}
