package screen

import (
	"github.com/GDGVIT/socrates/schema"
	"log"
	"encoding/json"
	"net/http"
	"bytes"
)

func (s *Screen) putUpdate() {
	client := &http.Client{}

	body, err := json.Marshal(s.config)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:" + s.port + "/update", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Screen) getConfig() {
	res, err := http.Get("http://localhost:" + s.port + "/view")
	if err != nil {
		log.Fatal(err)
	}

	var config *schema.Config
	err = json.NewDecoder(res.Body).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	s.config = config
}