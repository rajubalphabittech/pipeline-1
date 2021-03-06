package amazon

import (
	"fmt"

	"github.com/banzaicloud/pipeline/pkg/providers/amazon"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Migrate executes the table migrations for the provider.
func Migrate(db *gorm.DB, logger logrus.FieldLogger) error {
	tables := []interface{}{
		&ObjectStoreBucketModel{},
	}

	var tableNames string
	for _, table := range tables {
		tableNames += fmt.Sprintf(" %s", db.NewScope(table).TableName())
	}

	logger.WithFields(logrus.Fields{
		"provider":    amazon.Provider,
		"table_names": tableNames,
	}).Info("migrating provider tables")

	return db.AutoMigrate(tables...).Error
}
