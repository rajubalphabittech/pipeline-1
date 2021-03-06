package providers

import (
	"github.com/banzaicloud/pipeline/internal/providers/amazon"
	"github.com/banzaicloud/pipeline/internal/providers/azure"
	"github.com/banzaicloud/pipeline/internal/providers/google"
	"github.com/banzaicloud/pipeline/internal/providers/oracle"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Migrate runs migrations for cloud provider services.
func Migrate(db *gorm.DB, logger logrus.FieldLogger) error {
	if err := amazon.Migrate(db, logger); err != nil {
		return err
	}

	if err := azure.Migrate(db, logger); err != nil {
		return err
	}

	if err := google.Migrate(db, logger); err != nil {
		return err
	}

	if err := oracle.Migrate(db, logger); err != nil {
		return err
	}

	return nil
}
