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

	handlerBook "github.com/vkhoa145/go-training/app/modules/books/handlers"
	repositoryBook "github.com/vkhoa145/go-training/app/modules/books/repositories"
	bookUseCase "github.com/vkhoa145/go-training/app/modules/books/usecase"
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

	// Category
	categoryRepo := repositoryCategory.NewCategoryRepo(server.DB)
	categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := handlerCategory.NewCategoryHandlers(categoryUseCase, categoryRepo)

	api.Post("/categories", authMiddleware, categoryHandler.CreateCategory())
	api.Get("/categories", authMiddleware, categoryHandler.GetAllCategories())
	api.Get("/categories/:id", authMiddleware, categoryHandler.GetCategoryById())
	api.Put("/categories/:id", authMiddleware, categoryHandler.UpdateCategory())
	api.Delete("/categories/:id", authMiddleware, categoryHandler.DeleteCategory())

	// Book
	bookRepo := repositoryBook.NewBookRepo(server.DB)
	bookUseCase := bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := handlerBook.NewBookHandlers(bookUseCase, bookRepo)

	api.Post("/categories/:id/books", authMiddleware, bookHandler.CreateBook())
	api.Get("/books/:id", bookHandler.GetBookById())
	api.Put("/categories/:id/books/:book_id", authMiddleware, bookHandler.UpdateBook())
	api.Delete("/books/:id", authMiddleware, bookHandler.DeleteBook())
}
