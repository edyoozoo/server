// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

// KVHistoryRequest
type KVHistoryRequest struct {
	Service_id int `json:"service_id"`
	Key_id     int `json:"key_id"`
}

func (r *KVHistoryRequest) GetService_id() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Service_id
}

func (r *KVHistoryRequest) GetKey_id() int {
	if r == nil {
		var zeroVal int
		return zeroVal
	}
	return r.Key_id
}
