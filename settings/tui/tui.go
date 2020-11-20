package main

import (
	"encoding/json"
	"log"
	"os"
	"github.com/rivo/tview"
	"tui/screen"
	"github.com/GDGVIT/socrates/schema"
	"net/http"
	"github.com/spf13/viper"
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
	// Setup logging
	file, err := os.OpenFile("socrates-client-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

	log.SetOutput(file)

	// Read config
	initConfig()
	
	// Get current settings
	port := viper.GetString("API_PORT")
	res, err := http.Get("http://localhost:" + port + "/view")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal response
	var config schema.Config
	err = json.NewDecoder(res.Body).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Start application
	app := tview.NewApplication()
	s := screen.New(app, &config, port)
	app.SetInputCapture(s.HandleInput)
	s.Start()
}