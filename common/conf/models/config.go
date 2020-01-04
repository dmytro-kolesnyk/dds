package models

// Describes all possible configuration options
type Config struct {
	CliApi struct {
		Port int `yaml:"port",envconfig:"CLIAPI_PORT"`
	} `yaml:"cliapi"`
	Storage struct {
		LocalStoragePath string `yaml:"local-storage-path",envconfig:"LOCAL_STORAGE_PATH"`
		DefaultStrategy  string `yaml:"strategy"`
		StoreLocal       bool   `yaml:"store-local"`
	} `yaml:"storage"`
	CommunicationServer struct {
		Port int `yaml:"target-port"`
	} `yaml:"communication-server"`
}
