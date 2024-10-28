package drivers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	gormPsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	SSLMode  string
	Timezone string
	client   *gorm.DB
	db       *sql.DB
}

// Connect establishes new connection to database.
func (postgres *Postgres) Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		postgres.Host, postgres.Port, postgres.Username,
		postgres.Password, postgres.Database, postgres.SSLMode, postgres.Timezone,
	)

	postgres.client, err = gorm.Open(gormPsql.Open(dsn))
	if err != nil {
		return err
	}

	postgres.db, _ = postgres.client.DB()

	return nil
}

// Close closes the connection to database.
func (postgres *Postgres) Close() error {
	return postgres.db.Close()
}

// GetClient returns an instance of database.
func (postgres *Postgres) GetClient() *gorm.DB {
	return postgres.client
}

// GetDB returns an instance of database.
func (postgres *Postgres) GetDB() *sql.DB {
	return postgres.db
}
