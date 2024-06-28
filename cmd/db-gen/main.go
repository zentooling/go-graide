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
			SingularTable: true, // singular table names (student vs students,etc)
		},
	})
	if err != nil {
		log.Fatalf("unable to connect to database: +%v\n", err)
	}

	// Migrate the schema
	if nil != db.AutoMigrate(&database.Institution{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Instructer{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Class{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Student{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Assignment{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Rubric{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Grade{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if nil != db.AutoMigrate(&database.Work{}) {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
}
