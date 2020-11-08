package model

import (
	// "log"
	"path"
	"github.com/GDGVIT/socrates/schema"
)

func loadConfig(path string) *schema.Config {
	config, err := schema.Load(path)

	if err != nil {
		config = schema.New()
		return config
	}
	return config
}

type Model struct {
	config	*schema.Config
	filePath	string
}

func New(folderPath string) *Model {
	filePath := path.Join(folderPath, "config.json")

	config := loadConfig(filePath)
	
	return &Model{
		config,
		filePath,
	}
}

func (m *Model) GetConfig() *schema.Config {
	return m.config
}

func (m *Model) PutConfig(topics []string, freq int) {
	m.config.Set(topics, freq)
	m.config.Save(m.filePath)
}

func (m *Model) Debug() string {
	out := ""
	out += string(m.config.Freq)
	for _, topic := range m.config.Topics {
		out += " " + topic
	}
	return out
}