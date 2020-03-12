package pkg

import (
	"encoding/json"
)

//Serializer is a function that takes one type and converts
//it into another type
func Serializer(in interface{}) ([]byte, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	return b, nil
}