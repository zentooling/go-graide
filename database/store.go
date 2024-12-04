package database

import (
	"fmt"

	"github.com/zentooling/graide/internal/config"
	"github.com/zentooling/graide/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var log = logger.New("schema-gen")

var DB *gorm.DB = nil

func Initialize() {
	cfg := config.Instance()

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

	DB = db
}

func Shutdown() {
	dbInstance, _ := DB.DB()
	_ = dbInstance.Close()
}

type InstitutionStore struct{}

func (store InstitutionStore) GetAll() []Institution {
	var ret []Institution

	DB.Find(&ret)

	return ret
}

type StudentStore struct{}

func (store StudentStore) GetById(id uint) *Student {
	ret := &Student{}

	if db := DB.First(ret, id); db.Error != nil {
		log.Printf("unable to find student with id %v: %v", id, db.Error)
		return nil
	}

	return ret
}

func (store StudentStore) GetByIdWithClasses(id uint) *Student {
	student := Student{}

	if db := DB.Model(&student).Preload("Classes").Find(&student, id); db.Error != nil {
		log.Printf("unable to find classes with student with id %v: %v", id, db.Error)
		return nil
	}
	return &student
}
