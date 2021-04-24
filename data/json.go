package data

import (
	"encoding/json"
	"io"
)	

func (d *RegisterPostSuccess) RegisterPostSuccessToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *RegisterUserStructure) FromJSONToRegisterUserStructure (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}