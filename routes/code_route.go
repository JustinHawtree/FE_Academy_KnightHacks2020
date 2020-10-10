package routes

import (
	"FoundationHelper_KnightHacks2020/app"
	"FoundationHelper_KnightHacks2020/controller"
)

// CodeRoutes handles routing for code api
func CodeRoutes() {
	codeRoute := app.App.Group("/code")

	codeRoute.Post("", controller.RunCode)
	codeRoute.Get("", controller.GetCode)
}
