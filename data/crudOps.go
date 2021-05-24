package data

func RegisterUserCRUDOPS(d *RegisterUserStructure) *RegisterPostSuccess{
	var response RegisterPostSuccess
	
	_,_ = UpdateUserToDatabase(d)
	
	response = RegisterPostSuccess{
		UserID:d.UserID,
		Message:"Successfully updated user details!",
		Status:200,
	}

	return &response
}		

func RegisterUserToBusinessCRUDOPS (d *RegisterUserToBusinessStruct) *RegisterToBusinessPostSuccess{
	var response RegisterToBusinessPostSuccess

	isSubscribed,isSubscribedErr := IsSubscribedAlready(d)
	if isSubscribedErr!=nil{
		response = RegisterToBusinessPostSuccess{
			BusinessID:d.BusinessID,
			Message:"Error! User not registered!",
		}
		return &response
	}

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

func GetUserIDCRUDOPS(phone string) *RegisterPostSuccess{
	var response RegisterPostSuccess
	var emptyString string
	var emptyData IdOfDoc
	if len(phone)<10 {
		response = RegisterPostSuccess{
			UserID: emptyString,
			Data: emptyData,
			Message:"Error! Less than 10 digit number!",
			Status: 401,
		}
	} else if len(phone)>10 {
		response = RegisterPostSuccess{
			UserID: emptyString,
			Data: emptyData,
			Message:"Error! More than 10 digit number!",
			Status: 401,
		}
	} else {
		userID,err := GetUserIDByPhoneMongo(phone)
		if err != nil {

			//Create new user with phone number
			payload := &RegisterUserStructure{
				UserName: emptyString,
				PhoneNumber: phone,
			}
			uid,_ := AddUserToDatabase(payload)
			
			response = RegisterPostSuccess{
				UserID: uid,
				Data: emptyData,
				Message:"successfully created new UserID",
				Status: 200,
			} 	
		} else{
			response = RegisterPostSuccess{
				UserID:userID.ID.Hex(),
				Data: userID,
				Message:"Success! Existing user",
				Status: 200,
			} 
		}
	}
	
	return &response
}	
