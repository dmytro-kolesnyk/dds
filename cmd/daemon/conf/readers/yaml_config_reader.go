package readers

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"gopkg.in/yaml.v2"
	"os"
)

// Reads configuration properties from specified YAML file
type YamlConfigReader struct {
	path string

	ConfigReader
}

func (rcv *YamlConfigReader) Read() (*models.Config, error) {
	file, err := os.Open(rcv.path)
	if err != nil {
		return nil, err
	}

	config := &models.Config{}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewYamlConfigReader(path string) *YamlConfigReader {
	return &YamlConfigReader{path: path}
}
