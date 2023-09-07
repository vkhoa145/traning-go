package server

import (
	_ "fmt"

	"github.com/vkhoa145/go-training/app/middlewares"
	handlerUser "github.com/vkhoa145/go-training/app/modules/users/handlers"
	repositoryUser "github.com/vkhoa145/go-training/app/modules/users/repositories"
	userUseCase "github.com/vkhoa145/go-training/app/modules/users/usecase"

	handlerCategory "github.com/vkhoa145/go-training/app/modules/categories/handlers"
	repositoryCategory "github.com/vkhoa145/go-training/app/modules/categories/repositories"
	categoryUseCase "github.com/vkhoa145/go-training/app/modules/categories/usecase"
)

func SetupRoutes(server *Server) {
	// Auth
	userRepo := repositoryUser.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := handlerUser.NewUserHandlers(userUseCase)

	authMiddleware := middlewares.NewAuthMiddleware(server.Config.SIGNED_STRING)

	api := server.Fiber.Group("/api/v1")
	user := api.Group("/users")
	user.Post("/signup", userHandler.SignUpUser(server.Config))
	user.Post("/signin", userHandler.SignInUser(server.Config))
	user.Get("/profile", authMiddleware, userHandler.GetUser())

	// Category
	categoryRepo := repositoryCategory.NewCategoryRepo(server.DB)
	categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := handlerCategory.NewCategoryHandlers(categoryUseCase)

	categories := api.Group("/categories")
	categories.Post("/create", authMiddleware, categoryHandler.CreateCategory())
}
