package database

func CreateInstitution() *Institution {
	return &Institution{
		Name:        "Denver Public Schools",
		Street:      "1233 South Broadway St.",
		Street2:     "",
		City:        "Denver",
		State:       "CO",
		Zip:         "80010",
		Phone:       "303 555-1212",
		Email:       "info@dps.edu",
		Instructors: make([]Instructor, 0),
	}
}
func CreateInstructor(institution *Institution) *Instructor {
	return &Instructor{
		InstitutionID: institution.ID,
		FirstName:     "Paula",
		LastName:      "Zendle",
		Phone:         "303 555-1212",
		Email:         "pzendle@dps.edu",
		Classes:       make([]Class, 0),
	}
}
func CreateClass(instructor *Instructor, name string, desc string) *Class {
	return &Class{
		InstructorID: instructor.ID,
		Name:         name,
		Description:  desc,
		Students:     nil,
		Assignments:  nil,
		Works:        nil,
	}
}
func CreateStudent(firstName string, lastName string) *Student {
	return &Student{
		Classes:   nil,
		Works:     nil,
		FirstName: firstName,
		LastName:  lastName,
	}
}
