package forms

import (
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}


func (f *Form) Required(fields ...string){
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == ""{
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == ""{
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}