package database

import (
	"fmt"
	"log"
	"os"
	"time"
	"toester/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	host     string
	user     string
	password string
	dbname   string
	port     uint
	sslmode  bool
}

func _default() config {
	return config{
		host:     "localhost",
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		port:     5432,
		sslmode:  false,
	}
}

func (c config) build() string {
	return fmt.Sprintf(
		`host=%s
     port=%d 
     user=%s 
     password=%s 
     dbname=%s`,
		c.host,
		c.port,
		c.user,
		c.password,
		c.dbname,
	)
}

func Must() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	_ = newLogger
	db, err := gorm.Open(postgres.Open(_default().build()), &gorm.Config{})
	{
		if err != nil {
			panic(err)
		}
	}

	if err := db.AutoMigrate(model.PlayerModel{}, model.UserModel{}); err != nil {
		panic(err)
	}

	return db
}
