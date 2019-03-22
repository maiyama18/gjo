package encode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const indent = "  "

func Encode(args []string, array, pretty bool) (string, error) {
	if array {
		return encodeArray(args, pretty)
	} else {
		return encodeObject(args, pretty)
	}
}

func encodeObject(args []string, pretty bool) (string, error) {
	keyValues := make(map[string]interface{})
	for _, arg := range args {
		key, value, err := parseKeyValue(arg)
		if err != nil {
			continue
		}
		keyValues[key] = value
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if pretty {
		enc.SetIndent("", indent)
	}

	err := enc.Encode(keyValues)
	return buf.String(), err
}

func encodeArray(args []string, pretty bool) (string, error) {
	values := make([]interface{}, 0)
	for _, arg := range args {
		value, err := parseValue(arg)
		if err != nil {
			continue
		}
		values = append(values, value)
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if pretty {
		enc.SetIndent("", indent)
	}

	err := enc.Encode(values)
	return buf.String(), err
}

func parseKeyValue(arg string) (string, interface{}, error) {
	splitted := strings.Split(arg, "=")
	if len(splitted) != 2 {
		return "", "", fmt.Errorf("parse error: %s", arg)
	}

	val, err := parseValue(splitted[1])
	if err != nil {
		return "", "", err
	}
	return splitted[0], val, nil
}

func parseValue(arg string) (interface{}, error) {
	if val, err := strconv.Atoi(arg); err == nil {
		return val, nil
	}
	if val, err := strconv.ParseBool(arg); err == nil {
		return val, nil
	}
	return arg, nil
}
