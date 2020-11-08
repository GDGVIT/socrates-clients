package routes

import (
	"api/controller"
	"github.com/julienschmidt/httprouter"
)

// MakeRoutes registers routes on the router
func MakeRoutes(rtr *httprouter.Router, ctrl *controller.Controller) {
	rtr.GET("/view", ctrl.Viewer.View)
	rtr.PUT("/update", ctrl.Updater.Update)
}