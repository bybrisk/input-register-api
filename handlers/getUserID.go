package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/input-register-api/data"
)

// swagger:route GET /user/{PhoneNumber} user getUserIDByPhone
// Get userID or create UserID if not registered.
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

	w.Header().Set("Content-Type", "application/json")
	
	err := lp.RegisterPostSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}