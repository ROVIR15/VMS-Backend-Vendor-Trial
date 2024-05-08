package infrastructure

import (
	"database/sql"

	"emperror.dev/errors"
	db "github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/db"
)

func (ic *infrastructureConfigurator) configDatabase() (*sql.DB, func(), error) {
	db, err := db.NewDatabaseConnection(ic.config.PgOptions)
	if err != nil {
		return nil, nil, errors.WrapIf(err, "db.NewDatabaseConnection")
	}

	ic.logger.Infof("service's database established")

	return db.SqlDB, func() {
		db.SqlDB.Close()
		ic.logger.Infof("service's database disconected")
	}, nil
}
