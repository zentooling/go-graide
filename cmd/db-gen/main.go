package main

import (
	"fmt"

	"github.com/zentooling/graide/database"
	"github.com/zentooling/graide/internal/config"
	"github.com/zentooling/graide/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {

	var log = logger.New("schema-gen")
	var cfg = config.New("config.yml")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Denver",
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Database,
		cfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,     // singular table names (student vs students,etc)
		},
	})
	if err != nil {
		log.Fatalf("unable to connect to database: +%v\n", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&database.Institution{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Instructer{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Class{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Student{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Assignment{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Rubric{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Grade{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	err = db.AutoMigrate(&database.Work{})
	if err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
}
