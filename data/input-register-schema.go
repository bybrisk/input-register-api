package data

import (
	"github.com/go-playground/validator/v10"
)

//post request for registering a user
type RegisterUserStructure struct{
	// UserID of the user (system generated. No need to provide)
	//
	// required: false
	// max length: 1000
	UserID string `json: "-"`

	// The full Name of the user
	//
	// required: true
	// max length: 1000
	UserName string `json: "userName" validate:"required"`

	//Phone number of the customer
	//
	// required: true
	// max length: 1000
	PhoneNumber string `json: "phoneNumer" validate:"required"`

	// Complete address of the user
	//
	// required: true
	// max length: 1000
	Address string `json: "address" validate:"required"`

	// Latitude of the customer
	//
	// required: true
	// max length: 1000
	Latitude float64 `json: "latitude" validate:"required"`

	// Longitude of the customer
	//
	// required: true
	// max length: 1000
	Longitude float64 `json: "longitude" validate:"required"`
}

//post request for registering a user
type RegisterUserToBusinessStruct struct{
	// UserID of the user 
	//
	// required: true
	// max length: 1000
	UserID string `json: "userID" validate:"required"`

	// BusinessID of the business user is subscribing to
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessID" validate:"required"`
}

//post response
type RegisterPostSuccess struct {
	//userID of the user
	//
	UserID string `json:"userID"`
	//Message response
	//
	Message string `json:"message"`
}

//post response
type RegisterToBusinessPostSuccess struct {
	//businessID of the business
	//
	BusinessID string `json:"businessID"`
	//Message response
	//
	Message string `json:"message"`
}

func (d *RegisterUserStructure) ValidateRegisterUserStructure() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *RegisterUserToBusinessStruct) ValidateRegisterUserToBusinessStruct() error {
	validate := validator.New()
	return validate.Struct(d)
}