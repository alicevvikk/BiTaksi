package data

import "encoding/json"

import "io"

type MatchingRequest struct{
	Type		string	   `json:"type"`
	Coordinates	[]float64  `json:"coordinates"`
}

func (mr *MatchingRequest) ToJSON(w io.Writer) error {
	newEncoder := json.NewEncoder(w)
	return newEncoder.Encode(mr)
}

func (mr *MatchingRequest) FromJSON(r io.Reader) error {
	newDecoder := json.NewDecoder(r)
	return newDecoder.Decode(mr)
}

