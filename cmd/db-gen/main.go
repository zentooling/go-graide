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
	if err = db.AutoMigrate(&database.Institution{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Instructor{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Class{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Student{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Assignment{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Rubric{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Grade{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}
	if err = db.AutoMigrate(&database.Work{}); err != nil {
		log.Fatalf("AutoMigrate: +%v\n", err)
	}

	// create reference data

	inst := database.CreateInstitution()
	db.Create(&inst)
}
