package view

import (
	"log"
	"encoding/json"
	"api/model"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	configModel *model.Model
}

func (ctrl *Controller) View(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config := ctrl.configModel.GetConfig()
	res, err := json.Marshal(config)
	fmt.Printf("%+v\n%+v\n", string(res), config)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

func New(configModel *model.Model) *Controller {
	ctrl := Controller{configModel}
	return &ctrl
}
