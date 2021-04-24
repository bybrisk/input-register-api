
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-register-api/data"
)

// swagger:route POST /input/register/business user registerToBusiness
// Subscribe a user to a business.
//
// responses:
//	200: registerToBusinessPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Register) Subscribe_User (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-register-api Module")
	registeration := &data.RegisterUserToBusinessStruct{}

	err:=registeration.FromJSONToRegisterUserToBusinessStruct(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = registeration.ValidateRegisterUserToBusinessStruct()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-register-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add data to mongo
	response := data.RegisterUserToBusinessCRUDOPS(registeration)

	//writing to the io.Writer
	err = response.RegisterToBusinessPostSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}