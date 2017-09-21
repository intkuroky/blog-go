package main

import (
	"github.com/fpay/gopress"
	"github.com/intkuroky/blog-go/controllers"
	"github.com/intkuroky/blog-go/middlewares"
	"github.com/intkuroky/blog-go/errors"
	"github.com/tylerb/graceful"
	"time"
)

func main() {
	// create server
	s := gopress.NewServer(gopress.ServerOptions{
		Port: 3000,
	})

	s.App().HideBanner = true
	s.App().Debug = true

	errors.Debug = s.App().Debug
	s.App().HTTPErrorHandler = errors.CustomHTTPErrorHandler

	// init and register services
	// s.RegisterServices(
	// 	services.NewDatabaseService(),
	// )

	// register middlewares
	s.RegisterGlobalMiddlewares(
		gopress.NewLoggingMiddleware("global", nil),
	)

	s.App().Group("/articles", middlewares.NewJwtMiddleware())
	s.App().Group("/users/center", middlewares.NewJwtMiddleware())

	// init and register controllers
	s.RegisterControllers(
		controllers.NewUsersController(),
	)

	go s.Start()

	graceful.ListenAndServe(s.App().Server, 5*time.Second)
}
