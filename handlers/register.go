
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-register-api/data"
)

// swagger:route POST /input/register/create input-register registerAUser
// Register a user to input tool.
//
// responses:
//	200: registerPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) Register_User (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-register-api Module")
	registeration := &data.RegisterUserStructure{}

	err:=registeration.FromJSONToRegisterUserStructure(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = registeration.ValidateRegisterUserStructure()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-register-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add data to mongo
	response := data.RegisterUserCRUDOPS(registeration)

	//writing to the io.Writer
	err = response.RegisterPostSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}