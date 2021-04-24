package data

func RegisterUserCRUDOPS(d *RegisterUserStructure) *RegisterPostSuccess{
	var response RegisterPostSuccess
	
	id := AddUserToDatabase(d)
	
	response = RegisterPostSuccess{
		UserID:id,
		Message:"User registered successfully!",
	}

	return &response
}		