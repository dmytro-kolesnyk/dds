package readers

import (
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
)

// Describes reading application configuration
type ConfigReader interface {
	Read() (*models.Config, error)
}
