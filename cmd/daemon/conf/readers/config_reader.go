package readers

import "github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"

// Describes reading application configuration
type ConfigReader interface {
	Read() (*models.Config, error)
}
