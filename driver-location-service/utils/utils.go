package utils

import (
	"io"
	"encoding/json"
)

type locationType interface{}

// This util function encodes a 'locationType' into JSON and
// writes it into w.
func ToJSON (w io.Writer, l locationType) error {
	newEncoder := json.NewEncoder(w)
        return newEncoder.Encode(l)
}

// This util function reads from a 'io.Reader'
// and decodes the JSON value into given 'locationType'.
func FromJSON(r io.Reader, l locationType) error {
        newDecoder := json.NewDecoder(r)
        return newDecoder.Decode(l)
}

