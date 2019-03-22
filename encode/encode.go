package encode

import (
	"bytes"
	"fmt"
)

func Encode(args []string, array, pretty bool) (string, error) {
	if array {
		return encodeArray(args, pretty)
	} else {
		return encodeObject(args, pretty)
	}
}

func encodeObject(args []string, pretty bool) (string, error) {
	buf := &bytes.Buffer{}
	fmt.Println("object")
	fmt.Printf("%+v\n", args)
	return buf.String(), nil
}

func encodeArray(args []string, pretty bool) (string, error) {
	buf := &bytes.Buffer{}
	fmt.Println("array")
	fmt.Printf("%+v\n", args)
	return buf.String(), nil
}
