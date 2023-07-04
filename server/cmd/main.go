package main

import (
	"vue-fiber/config"
	"vue-fiber/internal/user"
	"vue-fiber/supa"
	"vue-fiber/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func main() {
	logger, err := util.InitLogger()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer util.CloseLogger()

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not initialise environment")
	}

	supa.Init(cfg)

	userSvc := user.NewService(supa.Client, logger)
	userHandler := user.NewHandler(userSvc, logger)

	r := fiber.New(fiber.Config{
		Prefork: true,
	})

	r.Static("/", "../client/dist")
	r.Use(recover.New())
	r.Use(cors.New(cors.Config{
		Next:         nil,
		AllowOrigins: "http://localhost:3030",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.Post("/login", userHandler.Login)
	auth.Post("/signup", userHandler.CreateUser)
	auth.Post("/logout", userHandler.Logout)

	if err := r.Listen(":3000"); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start the router")
	}
}
