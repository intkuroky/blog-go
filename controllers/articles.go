package controllers

import (
	"net/http"

	"github.com/fpay/gopress"
)

// ArticlesController
type ArticlesController struct {
	// Uncomment this line if you want to use services in the app
	// app *gopress.App
}

// NewArticlesController returns articles controller instance.
func NewArticlesController() *ArticlesController {
	return new(ArticlesController)
}

// RegisterRoutes registes routes to app
// It is used to implements gopress.Controller.
func (c *ArticlesController) RegisterRoutes(app *gopress.App) {
	// Uncomment this line if you want to use services in the app
	// c.app = app

	app.GET("/articles/sample", c.SampleGetAction)
	// app.POST("/articles/sample", c.SamplePostAction)
	// app.PUT("/articles/sample", c.SamplePutAction)
	// app.DELETE("/articles/sample", c.SampleDeleteAction)
}

// SampleGetAction Action
// Parameter gopress.Context is just alias of echo.Context
func (c *ArticlesController) SampleGetAction(ctx gopress.Context) error {
	// Or you can get app from request context
	// app := gopress.AppFromContext(ctx)
	data := map[string]interface{}{}
	return ctx.Render(http.StatusOK, "articles/sample", data)
}
