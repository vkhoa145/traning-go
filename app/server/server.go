package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vkhoa145/go-training/config"
	"github.com/vkhoa145/go-training/db"
	"gorm.io/gorm"
)

// definition Server struct
type Server struct {
	Fiber  *fiber.App
	DB     *gorm.DB
	Config *config.Config
}

// New Server function
func NewServer(cfg *config.Config) *Server {
	return &Server{
		Fiber:  fiber.New(),
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (server *Server) Start() error {
	server.Fiber.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	SetupRoutes(server)
	return server.Fiber.Listen(":" + server.Config.HTTP.Port)
}
