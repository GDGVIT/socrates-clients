package update

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"api/model"
	"github.com/GDGVIT/socrates/schema"
)

type Controller struct {
	configModel *model.Model
}

func (ctrl *Controller) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	config := &schema.Config{}
	fmt.Printf("In update\n%+v\n", r.Body)
	if err := json.NewDecoder(r.Body).Decode(config); err != nil {
		log.Println(err)
		return
	}
	ctrl.configModel.PutConfig(config.Topics, config.Freq)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Config successfully updated")
}

func New(configModel *model.Model) *Controller {
	ctrl := Controller{configModel}
	return &ctrl
}
