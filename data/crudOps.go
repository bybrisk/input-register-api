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

func RegisterUserToBusinessCRUDOPS (d *RegisterUserToBusinessStruct) *RegisterToBusinessPostSuccess{
	var response RegisterToBusinessPostSuccess

	isSubscribed := IsSubscribedAlready(d)

	if (isSubscribed==0){
		nameOfBusiness,BusinessCategory := GetBusinessName(d.BusinessID)
		if (nameOfBusiness!="" && BusinessCategory!="") {
			_ = RegisterToBusinessMongo(d,nameOfBusiness,BusinessCategory)
			response = RegisterToBusinessPostSuccess{
				BusinessID:d.BusinessID,
				Message:"Success! User subscribed successfully!",
			}
		} else{
			response = RegisterToBusinessPostSuccess{
				BusinessID:d.BusinessID,
				Message:"Error! User subscription failed!",
			}
		}
	} else{
		response = RegisterToBusinessPostSuccess{
			BusinessID:d.BusinessID,
			Message:"Warning! User already subscribed!",
		}
	}

	return &response
}