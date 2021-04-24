// Package classification of Input-Register API
//
// Documentation for Input-Register API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/input-register-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Success message on a single user registeration
// swagger:response registerPostResponse
type accountPostResponseWrapper struct {
	// Success message on newly registered user
	// in: body
	Body data.RegisterPostSuccess
}

// swagger:parameters registerAUser
type registerAUserParmsWrapper struct {
	// Data structure to register a user.
	// in: body
	// required: true
	Body data.RegisterUserStructure
}

// swagger:parameters registerToBusiness
type registerAUserToBusinessParmsWrapper struct {
	// Data structure to register a user to a business.
	// in: body
	// required: true
	Body data.RegisterUserToBusinessStruct
}

// Success message on subscribing to a business
// swagger:response registerToBusinessPostResponse
type accountPostResponseWrapper struct {
	// Success message on subscribing to a business
	// in: body
	Body data.RegisterToBusinessPostSuccess
}