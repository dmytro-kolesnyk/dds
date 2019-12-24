package conf

import (
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"github.com/dmytro-kolesnyk/dds/common/conf/readers"
	"github.com/imdario/mergo"
)

type Resolver struct {
}

// Returns conf object in order: YAML, ENV.
func (rcv *Resolver) GetConfig() (*models.Config, error) {
	envConfigReader := readers.NewEnvConfigReader()
	yamlConfigReader := readers.NewYamlConfigReader("config.yaml")

	return merge(yamlConfigReader, envConfigReader)
}

func merge(readers ...readers.ConfigReader) (*models.Config, error) {
	config := &models.Config{}

	for _, reader := range readers {
		newConfig, err := reader.Read()
		if err != nil {
			return nil, err
		}
		err = mergo.Merge(config, newConfig, mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func NewResolver() *Resolver {
	return &Resolver{}
}
