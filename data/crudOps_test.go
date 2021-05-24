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
		BusinessID: "606d6dd51bf7f6ed4a0b320c",
		UserID: "6083deb86fcd474489784fee",
	}

	res:= data.RegisterUserToBusinessCRUDOPS(register) 

	fmt.Println(res)
}*/

func TestGetUserIDCRUDOPS(t *testing.T){
	res:= data.GetUserIDCRUDOPS("9079528682")
	fmt.Println(res)
}