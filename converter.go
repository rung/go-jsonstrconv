package jsonstrconv

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
)

// ToString converts all types in your json to string type
// It doesn't support NDJSON now. Please use this for each json.
func ToString(payload []byte) ([]byte, error) {
	if !json.Valid(payload) {
		return nil, errors.New("Failed to parse json")
	}
	converted, err := converter(payload)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert")
	}

	return converted, nil
}

func converter(input []byte) ([]byte, error) {
	data := []rune(string(input))
	buf := bytes.Buffer{}

	for pos := 0; pos < len(data); pos++ {
		orgPos := pos

		switch {
		case data[pos] == '"':
			i, err := readString(data[pos:])
			if err != nil {
				return nil, err
			}
			pos += i
			buf.WriteString(string(data[orgPos : pos+1]))
		case isValue(data[pos]):
			pos += readValue(data[pos:])
			buf.WriteString("\"" + string(data[orgPos:pos+1]) + "\"")
		default:
			buf.WriteRune(data[pos])
		}
	}
	return buf.Bytes(), nil
}

func readString(data []rune) (int, error) {
	var escaped bool
	for i, v := range data[1:] {
		if escaped {
			escaped = false
			continue
		}
		switch v {
		case '\\':
			escaped = true
		case '"':
			return i + 1, nil
		}
	}
	return 0, errors.New("string doesn't have '\"'")
}

func readValue(data []rune) int {
	var i int
	var v rune
	for i, v = range data {
		switch v {
		case ' ', '\t', '\n', '\r', ',', '}', ']':
			return i - 1
		}
	}
	return i
}

func isValue(r rune) bool {
	switch r {
	// num, true, false, undefined or null
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', 't', 'f', 'u', 'n':
		return true
	}
	return false
}
