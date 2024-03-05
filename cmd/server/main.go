package main

// pdfRoutes "api-crud-golang/internal/modules/user/infra/routes/pdf"

//"github.com/gin-gonic/gin"

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	// Provide dependency injection
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	userHandler "api-crud-golang/internal/modules/user/infra/routes/handler"
	userService "api-crud-golang/internal/modules/user/services"

	userRepository "api-crud-golang/internal/modules/user/infra/orm/gorm/repository"
	middleware "api-crud-golang/internal/shared/infra/middleware/tokenCheck"
)

func NewHTTPServer(lc fx.Lifecycle) *gin.Engine {

	router := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{}

	srv := &http.Server{Addr: ":3333", Handler: router}

	//router.Use(cors.New(config))

	router.Use(middleware.JwtTokenCheck)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}  

	app := fx.New(
		fx.Provide(
			userService.NewServices,
			userRepository.NewUserRepository,
			 
			userHandler.NewHandler,
			NewHTTPServer,
		),
		fx.Invoke(
			userHandler.RoutesV1,
			func(router *gin.Engine) {},
		),
	)

	app.Run()
}
