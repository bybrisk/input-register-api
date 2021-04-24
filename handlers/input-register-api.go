package handlers

import (
	"log"
)

type Input_Register struct {
 l *log.Logger
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func New_Input_Register(l *log.Logger) *Input_Register{
	return &Input_Register{l}
}