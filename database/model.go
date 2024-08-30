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
	Instructors []Instructor
}

type Instructor struct {
	gorm.Model
	InstitutionID uint
	FirstName     string
	LastName      string
	Phone         string
	Email         string
	Classes       []Class
}
type Class struct {
	gorm.Model
	InstructorID uint
	Name         string
	Description  string
	Students     []Student `gorm:"many2many:class_x_student;"`
	Assignments  []Assignment
	Works        []Work
}
type Student struct {
	gorm.Model
	Classes   []Class `gorm:"many2many:class_x_student;"`
	Works     []Work
	FirstName string
	LastName  string
}

type Work struct {
	gorm.Model
	StudentID    uint
	ClassID      uint
	GradeID      uint
	AssignmentID uint
	Work         string
}

type Assignment struct {
	gorm.Model
	ClassID  uint
	RubricID uint
	Works    []Work
	Text     string
}

type Rubric struct {
	gorm.Model
	Assignments []Assignment
	Text        string
}

type Grade struct {
	gorm.Model
	Works []Work
	Grade string
}
