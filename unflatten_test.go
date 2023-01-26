package flat

import "testing"

func TestUnflatten(t *testing.T) {
	type test struct {
		Value string `json:"inner.value"`
	}

	data, err := Unflatten(test{Value: "hello"})
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != `{"inner":{"value":"hello"}}` {
		t.Fatal("unflatten failed. Expected: {\"inner\":{\"value\":\"hello\"}}. Got: ", string(data))
	}
}
