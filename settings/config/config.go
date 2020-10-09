package config

import (
	// "encoding/json"
)

type Config struct {
	topics []string 		`json: "topics"`
	freq int 				`json: "frequency`
}

func New() Config {
	return Config {
		nil,
		0,
	}
}