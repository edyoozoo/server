// Code generated by protoapi:go; DO NOT EDIT.

package searchsvr

// FieldError
type FieldError struct {
	FieldName string             `json:"fieldName"`
	ErrorType *ValidateErrorType `json:"errorType"`
}

func (r *FieldError) GetFieldName() string {
	if r == nil {
		var zeroVal string
		return zeroVal
	}
	return r.FieldName
}

func (r *FieldError) GetErrorType() *ValidateErrorType {
	if r == nil {
		var zeroVal *ValidateErrorType
		return zeroVal
	}
	return r.ErrorType
}
