package db

import (
	"database/sql"
	"fmt"
	"time"

	option "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/db/config"
	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseConn struct {
	SqlDB *sql.DB
}

func NewDatabaseConnection(config *option.PgOptions) (*DatabaseConn, error) {
	var url string
	switch config.DBDriver {
	case "postgres":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.DBHost,
			config.DBPort,
			config.DBUser,
			config.DBPassword,
			config.DBName,
			config.DBSSLMode,
		)
	default:
		return nil, fmt.Errorf("connection name '%v' is not supported", config.DBDriver)
	}

	// Open a new connection to the database.
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("error, not open connection to database, %w", err)
	}

	// Set the maximum number of open connections to the database.
	db.SetMaxOpenConns(config.DBMaxOpenConn)
	db.SetMaxIdleConns(config.DBMaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(config.DBMaxLifetimeConn) * time.Second)

	if config.Debug {
		boil.SetDB(db)
		boil.DebugMode = true
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return &DatabaseConn{SqlDB: db}, nil
}
