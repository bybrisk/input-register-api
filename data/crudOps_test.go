package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/input-register-api/data"
)

/*func TestRegisterUserCRUDOPS(t *testing.T) {

	register := &data.RegisterUserStructure{
		UserID: "60ab544b4306be85b5f35fba",
		UserName: "Shashank Prakash",
		PhoneNumber: "9340232345",
		Address: "Maulana Azad National Institute of Technology, Bhopal, MP",
		Latitude: 23.123456789,
		Longitude: 77.12345678,
	}

	res:= data.RegisterUserCRUDOPS(register) 

	fmt.Println(res)
	if res==nil{
		t.Fail()
	}
}*/

/*func TestRegisterUserCRUDOPS(t *testing.T) {

	register := &data.RegisterUserToBusinessStruct{
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
		UserID: "6083deb86fcd474489784fee",
	}

	res:= data.RegisterUserToBusinessCRUDOPS(register) 

	fmt.Println(res)
}*/

func TestGetUserIDCRUDOPS(t *testing.T){
	res:= data.GetUserIDCRUDOPS("9340232345")
	fmt.Println(res)
}

/*func TestRegisterUserToBusinessCRUDOPS(t *testing.T) {
	payload := &data.RegisterUserToBusinessStruct{
		UserID: "6083deb86fcd474489784fee",
		BusinessID: "6038bd0fc35e3b8e8bd9f81a",
	}

	res := data.RegisterUserToBusinessCRUDOPS(payload)
	fmt.Println(res)
}*/