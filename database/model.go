package database

import "gorm.io/gorm"

type Institution struct {
	gorm.Model
	Name        string
	Street      string
	Street2     string
	City        string
	State       string
	Zip         string
	Phone       string
	Email       string
	Instructers []Instructer
}

type Instructer struct {
	gorm.Model
	InstitutionID uint
	Name          string
	Classes       []Class
}
type Class struct {
	gorm.Model
	InstructerID uint
	Name         string
	Students     []Student `gorm:"many2many:class_x_student;"`
	Assignments  []Assignment
}
type Student struct {
	gorm.Model
	Classes     []Class `gorm:"many2many:class_x_student;"`
	Assignments []Assignment
	Name        string
}

type Work struct {
	StudentID    uint
	AssignmentID uint
	GradeID      uint
	Work         string
}

type Assignment struct {
	gorm.Model
	StudentID uint
	ClassID   uint
	Text      string
	RubricID  uint
}

type Rubric struct {
	gorm.Model
	Text string
}

type Grade struct {
	gorm.Model
	Value string
}
