package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/input-register-api/data"
)

// swagger:route GET /user/{PhoneNumber} user getUserIDByPhone
// Get userID by 10 digit phone number registered to the account.
//
// responses:
//	200: registerPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Register) GetUserIDByPhone (w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> user-api Module")
	
	vars := mux.Vars(r)
	num := vars["PhoneNumber"]

	lp := data.GetUserIDCRUDOPS(num)

	err := lp.RegisterPostSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}