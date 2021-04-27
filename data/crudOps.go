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

	if len(phone)<10 {
		response = RegisterPostSuccess{
			UserID:"",
			Message:"Error! Less than 10 digit number!",
		}
	} else if len(phone)>10 {
		response = RegisterPostSuccess{
			UserID:"",
			Message:"Error! More than 10 digit number!",
		}
	} else {
		userID,err := GetUserIDByPhoneMongo(phone)
		if err!=nil{
			response = RegisterPostSuccess{
				UserID:"",
				Message:err.Error(),
			} 	
		} else{
			response = RegisterPostSuccess{
				UserID:userID,
				Message:"Success",
			} 
		}
	}
	
	return &response
}	
