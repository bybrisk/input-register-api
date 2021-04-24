package data_test

import (
	"testing"
	"fmt"
	"github.com/bybrisk/input-register-api/data"
)

func TestRegisterUserCRUDOPS(t *testing.T) {

	register := &data.RegisterUserStructure{
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
}