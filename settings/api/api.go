package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"api/controller"
	"api/model"
	"api/routes"
	"log"
	"os"
	"path"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Setup Config
	initConfig()

	// Setup logging
	logsPath := path.Join(viper.GetString("API_LOG_DIR"), "socrates-api-logs.txt")
	file, err := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)

	// Start app
	router := httprouter.New()
	model := model.New(viper.GetString("API_MODEL_DIR"))
	ctrl := controller.New(model)
	routes.Register(router, ctrl)

	log.Fatal(http.ListenAndServe(":" + viper.GetString("API_PORT"), router))
}