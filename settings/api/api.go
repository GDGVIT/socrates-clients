package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"api/controller"
	"api/model"
	"api/routes"
	"log"
	"os"
	"github.com/joho/godotenv"
	"path"
)

func main() {
	// Import env
	env, err := godotenv.Read(path.Join("..", "settings.env"))
	if err != nil {
		log.Fatalln(err)
	}

	// Setup logging
	logsPath := path.Join(env["API_LOG_DIR"], "socrates-api-logs.txt")
	file, err := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)

	// Start app
	router := httprouter.New()
	model := model.New(env["API_MODEL_DIR"])
	ctrl := controller.New(model)
	routes.Register(router, ctrl)

	log.Fatal(http.ListenAndServe(":" + env["API_PORT"], router))
}