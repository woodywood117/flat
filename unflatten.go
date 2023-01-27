package flat

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Unflattener struct {
	delimiter string
}

func NewUnflattener(delimiter string) *Unflattener {
	return &Unflattener{delimiter: delimiter}
}

func (u *Unflattener) Delimiter(delimiter string) {
	u.delimiter = delimiter
}

// TODO: Implement unflatten
func (u *Unflattener) Unflatten(item any) ([]byte, error) {
	if reflect.TypeOf(item).Kind() != reflect.Struct {
		return nil, errors.New("unflatten expects a struct as input")
	}

	return json.Marshal(item)
}

type meta struct {
	value reflect.Value
	field reflect.StructField
	path  []string
}

// fetchMeta is a recursive function that generates a list of metadata objects for each field in the struct
func (u *Unflattener) fetchMeta(item any, currentPath []string) (metadata []meta, err error) {
	if reflect.TypeOf(item).Kind() != reflect.Struct {
		errstr := fmt.Sprintf("cannot get metadata from a non-struct (%s)", strings.Join(currentPath, u.delimiter))
		return nil, errors.New(errstr)
	}

	ValueOf := reflect.ValueOf(item)

	fields := ValueOf.Type().NumField()
	for i := 0; i < fields; i++ {
		field := ValueOf.Type().Field(i)

		// Get the json path from the json struct tag
		var JsonPath string
		{
			tag, ok := field.Tag.Lookup("json")
			if ok {
				split := strings.Split(tag, ",")
				JsonPath = split[0]
			}

			if JsonPath == "" {
				JsonPath = field.Name
			}
		}

		// If the field is a struct, recursively call fetchMeta
		// Otherwise, append the metadata to the list
		SplitJsonPath := strings.Split(JsonPath, u.delimiter)
		if field.Type.Kind() == reflect.Struct {
			subPath := append(currentPath, SplitJsonPath...)
			subMetadata, err := u.fetchMeta(ValueOf.Field(i).Interface(), subPath)
			if err != nil {
				return nil, err
			}

			metadata = append(metadata, subMetadata...)
		} else {
			metadata = append(metadata, meta{
				value: ValueOf.Field(i),
				field: field,
				path:  append(currentPath, SplitJsonPath...),
			})
		}
	}

	return
}
