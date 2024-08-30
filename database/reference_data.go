package database

func CreateInstitution() Institution {
	return Institution{
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
