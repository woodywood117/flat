package flat

import "encoding/json"

// TODO: Implement Unflatten
func Unflatten(item any) ([]byte, error) {
	return json.Marshal(item)
}
