package postgres

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	ErrDsnIsEmpty         = errors.New("postgres dsn is empty")
	ErrDatabaseConnection = errors.New("could not connect to database")
)

func NewPostgresDb() (*gorm.DB, error) {
	dsn := getDsn()

	if len(dsn) == 0 {
		return nil, ErrDsnIsEmpty
	}

	pgConfig := postgres.Config{
		DriverName:           "",
		DSN:                  dsn,
		PreferSimpleProtocol: false,
		WithoutReturning:     false,
		Conn:                 nil,
	}

	gConfig := gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,

		// AllowGlobalUpdate enables the batch updates
		AllowGlobalUpdate: true,
		QueryFields:       false,
		CreateBatchSize:   2000,
		TranslateError:    false,
		ClauseBuilders:    nil,
		ConnPool:          nil,
		Dialector:         nil,
		Plugins:           nil,
	}

	var db *gorm.DB
	var err error

	// retry to connect to database
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.New(pgConfig), &gConfig)
		if err != nil {
			log.Printf("Try %d couldn't connect to DB: %s", i, err.Error())
			time.Sleep(3 * time.Second)
			continue
		} else {
			log.Printf("Try %d connected to DB", i)
			break
		}
	}

	if err != nil {
		return nil, ErrDatabaseConnection
	}

	return db, nil
}

func getDsn() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_DB_NAME")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintf("host=%s user=%s password= %s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		host, username, password, name, port)
}
