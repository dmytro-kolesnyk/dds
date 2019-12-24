package readers

import (
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"github.com/kelseyhightower/envconfig"
)

// Reads configuration properties from OS Environment Variables
type EnvConfigReader struct {
	ConfigReader
}

func (rcv *EnvConfigReader) Read() (*models.Config, error) {
	config := &models.Config{}

	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func NewEnvConfigReader() *EnvConfigReader {
	return &EnvConfigReader{}
}
