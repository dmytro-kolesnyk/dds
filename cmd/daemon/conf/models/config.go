package models

// Describes all possible configuration options
type Config struct {
	CliApi struct {
		Port int `yaml:"port",envconfig:"CLIAPI_PORT"`
	} `yaml:"cliapi"`
}
