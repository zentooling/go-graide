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

	cleanDb(db)

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
	db.Create(inst)
	instructor := database.CreateInstructor(inst)
	db.Create(instructor)
	eng101 := database.CreateClass(instructor, "ENG101", "No comprende")
	db.Create(eng101)
	eng102 := database.CreateClass(instructor, "ENG201", "Poco comprende")
	db.Create(eng102)
	eng103 := database.CreateClass(instructor, "ENG301", "comprende")
	db.Create(eng103)

	student := database.CreateStudent("Jose", "Estudio")
	db.Create(student)
	if err = db.Model(student).Association("Classes").Append(eng101, eng102); err != nil {
		log.Fatalf("Model.Associate: +%v\n", err)
	}

	student = database.CreateStudent("Maria", "Illegal")
	db.Create(student)
	if err = db.Model(student).Association("Classes").Append(eng102, eng103); err != nil {
		log.Fatalf("Model.Associate: +%v\n", err)
	}

	student = database.CreateStudent("Wentworth", "Noworksohard")
	db.Create(student)
	if err = db.Model(student).Association("Classes").Append(eng101, eng102, eng103); err != nil {
		log.Fatalf("Model.Associate: +%v\n", err)
	}

	rubric := &database.Rubric{Text: "CORE Standard"}
	db.Create(rubric)
	assignment := &database.Assignment{
		RubricID: rubric.ID,
		ClassID:  eng101.ID,
		Text:     "Change air in tires",
	}
	db.Create(assignment)
	assignment2 := &database.Assignment{
		RubricID: rubric.ID,
		ClassID:  eng101.ID,
		Text:     "Arrange sock drawer",
	}
	db.Create(assignment2)

	grade := &database.Grade{
		Grade: "None",
	}
	db.Create(grade)

	work := &database.Work{
		Model:        gorm.Model{},
		StudentID:    student.ID,
		ClassID:      eng101.ID,
		GradeID:      grade.ID,
		AssignmentID: assignment2.ID,
		Work:         "this is the location of my document",
	}
	db.Create(work)

}

func cleanDb(db *gorm.DB) {
	db.Exec("DELETE FROM CLASS_X_STUDENT CASCADE")
	db.Exec("DELETE FROM GRADE CASCADE")
	db.Exec("DELETE FROM WORK CASCADE")
	db.Exec("DELETE FROM ASSIGNMENT CASCADE")
	db.Exec("DELETE FROM RUBRIC CASCADE")
	db.Exec("DELETE FROM CLASS CASCADE")
	db.Exec("DELETE FROM INSTRUCTOR CASCADE")
	db.Exec("DELETE FROM STUDENT CASCADE")
	db.Exec("DELETE FROM INSTITUTION CASCADE")

}
