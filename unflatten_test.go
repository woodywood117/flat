package flat

import (
	"testing"
)

func TestUnflatten(t *testing.T) {
	type test struct {
		Value string `json:"inner.value"`
	}

	u := NewUnflattener(".")

	data, err := u.Unflatten(test{Value: "hello"})
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != `{"inner":{"value":"hello"}}` {
		t.Fatal("unflatten failed. Expected: {\"inner\":{\"value\":\"hello\"}}. Got: ", string(data))
	}
}

func TestFetchMetaSimple(t *testing.T) {
	type test struct {
		IKey   string `json:"inner.key"`
		IValue string `json:"inner.value"`
		OKey   string `json:"outer.key"`
		OValue string `json:"outer.value"`
	}

	u := NewUnflattener(".")

	metadata, err := u.fetchMeta(test{"hello", "world", "hello", "world"}, []string{})
	if err != nil {
		t.Fatal(err)
	}

	if len(metadata) != 4 {
		t.Fatal("fetchMeta failed. Expected 4 metadata object. Got: ", len(metadata))
	}

	if metadata[0].path[0] != "inner" || metadata[0].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[1].path[0] != "inner" || metadata[1].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[2].path[0] != "outer" || metadata[2].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[3].path[0] != "outer" || metadata[3].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
}

func TestFetchMetaInnerStruct(t *testing.T) {
	type inner struct {
		IKey   string `json:"key"`
		IValue string `json:"value"`
	}
	type test struct {
		Inner  inner  `json:"inner"`
		OKey   string `json:"outer.key"`
		OValue string `json:"outer.value"`
	}

	u := NewUnflattener(".")

	metadata, err := u.fetchMeta(test{inner{"hello", "world"}, "hello", "world"}, []string{})
	if err != nil {
		t.Fatal(err)
	}

	if len(metadata) != 4 {
		t.Fatal("fetchMeta failed. Expected 4 metadata object. Got: ", len(metadata))
	}

	if metadata[0].path[0] != "inner" || metadata[0].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[1].path[0] != "inner" || metadata[1].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[2].path[0] != "outer" || metadata[2].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[3].path[0] != "outer" || metadata[3].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
}

func TestFetchMetaInnerStructPointer(t *testing.T) {
	type inner struct {
		IKey   string `json:"key"`
		IValue string `json:"value"`
	}
	type test struct {
		Inner  *inner `json:"inner"`
		OKey   string `json:"outer.key"`
		OValue string `json:"outer.value"`
	}

	u := NewUnflattener(".")

	metadata, err := u.fetchMeta(test{&inner{"hello", "world"}, "hello", "world"}, []string{})
	if err != nil {
		t.Fatal(err)
	}

	if len(metadata) != 4 {
		t.Fatal("fetchMeta failed. Expected 4 metadata object. Got: ", len(metadata))
	}

	if metadata[0].path[0] != "inner" || metadata[0].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[1].path[0] != "inner" || metadata[1].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[2].path[0] != "outer" || metadata[2].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[3].path[0] != "outer" || metadata[3].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
}

func TestFetchMetaInnerNestedStructAndPointer(t *testing.T) {
	type value struct {
		Test1 string `json:"value"`
		Test2 string `json:"value2"`
	}
	type inner struct {
		IKey   string `json:"key"`
		IValue *value `json:"value"`
	}
	type test struct {
		Inner  inner  `json:"inner"`
		OKey   string `json:"outer.key"`
		OValue string `json:"outer.value"`
	}

	u := NewUnflattener(".")

	metadata, err := u.fetchMeta(test{inner{"hello", &value{"world1", "world2"}}, "hello", "world"}, []string{})
	if err != nil {
		t.Fatal(err)
	}

	if len(metadata) != 5 {
		t.Fatal("fetchMeta failed. Expected 5 metadata object. Got: ", len(metadata))
	}

	if metadata[0].path[0] != "inner" || metadata[0].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[1].path[0] != "inner" || metadata[1].path[1] != "value" || metadata[1].path[2] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[2].path[0] != "inner" || metadata[2].path[1] != "value" || metadata[2].path[2] != "value2" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[3].path[0] != "outer" || metadata[3].path[1] != "key" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
	if metadata[4].path[0] != "outer" || metadata[4].path[1] != "value" {
		t.Fatal("fetchMeta failed. Expected path to be 'inner.value'. Got: ", metadata[0].path)
	}
}
