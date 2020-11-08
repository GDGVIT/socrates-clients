package controller

import (
	"api/model"
	"api/controller/update"
	"api/controller/view"
)

type Controller struct {
	configModel *model.Model
	Updater	*update.Controller
	Viewer	*view.Controller
}

func New(configModel *model.Model) *Controller {
	updater := update.New(configModel)
	viewer := view.New(configModel)
	return &Controller {
		configModel,
		updater,
		viewer,
	}
}
