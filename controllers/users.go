package controllers

import (
	"net/http"
	"github.com/fpay/gopress"
)

// UsersController
type UsersController struct {
	// Uncomment this line if you want to use services in the app
	// app *gopress.App
}

// NewUsersController returns users controller instance.
func NewUsersController() *UsersController {
	return new(UsersController)
}
// RegisterRoutes registes routes to app
// It is used to implements gopress.Controller.
func (c *UsersController) RegisterRoutes(app *gopress.App) {
	// Uncomment this line if you want to use services in the app
	// c.app = app

	app.GET("/users/login", c.LoginGetAction)
	app.GET("/users/register", c.RegisterGetAction)
	// app.POST("/users/sample", c.SamplePostAction)
	// app.PUT("/users/sample", c.SamplePutAction)
	// app.DELETE("/users/sample", c.SampleDeleteAction)
}

// LoginGetAction Action
// Parameter gopress.Context is just alias of echo.Context
func (c *UsersController) LoginGetAction(ctx gopress.Context) error {
	// Or you can get app from request context
	// app := gopress.AppFromContext(ctx)
	//data := map[string]interface{}{}
	return ctx.String(http.StatusOK,"users/login")
	//return ctx.Render(http.StatusOK, "users/login", data)
}

// RegisterGetAction Action
func (c *UsersController) RegisterGetAction(ctx gopress.Context) error {
	//return ctx.String(http.StatusOK, "users/register")
	return ctx.Render(http.StatusOK, "users/register", nil)
}
