// Code generated by protoapi:go; DO NOT EDIT.

package searchsvr

// ValidateError
type ValidateError struct {
	Errors []*FieldError `json:"errors"`
}

func (r *ValidateError) GetErrors() []*FieldError {
	if r == nil {
		var zeroVal []*FieldError
		return zeroVal
	}
	return r.Errors
}
