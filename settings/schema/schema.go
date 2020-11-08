package schema

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Topics []string 		`json:"topics"`
	Freq int 				`json:"frequency"`
}

func New() *Config {
	return &Config {
		nil,
		0,
	}
}

func (c *Config) Save(path string) error {
	out, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}

	return ioutil.WriteFile(path, out, 0600)
}

// Load config from JSON file. Returns error if file not found, calls os.Exit if file cannot be parsed 
func Load(path string) (*Config, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var config Config 
	err = json.Unmarshal(contents, &config)
	if err != nil {
		log.Fatalln(err)
	}
	
	return &config, nil
}

// Set a config instance according to a new list of topics and frequency
func (c *Config) Set(topics []string, freq int) {
	c.Topics = topics
	c.Freq = freq
}
