package routes

import (
	"api/controller"
	"github.com/julienschmidt/httprouter"
)

func Register(rtr *httprouter.Router, ctrl *controller.Controller) {
	rtr.GET("/view", ctrl.Viewer.View)
	rtr.PUT("/update", ctrl.Updater.Update)
}