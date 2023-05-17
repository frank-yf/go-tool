//go:build jsoniter

package cjson

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	_jsonAPI            Interface = &jsonInterator{}
	MarshalToString               = _jsonAPI.MarshalToString
	Marshal                       = _jsonAPI.Marshal
	MarshalIndent                 = _jsonAPI.MarshalIndent
	UnmarshalFromString           = _jsonAPI.UnmarshalFromString
	Unmarshal                     = _jsonAPI.Unmarshal
	NewEncoder                    = _jsonAPI.NewEncoder
	NewDecoder                    = _jsonAPI.NewDecoder
	Valid                         = _jsonAPI.Valid

	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type jsonInterator struct{}

func (jsonInterator) MarshalToString(v any) (string, error) {
	return json.MarshalToString(v)
}

func (jsonInterator) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonInterator) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (jsonInterator) UnmarshalFromString(str string, v any) error {
	return json.UnmarshalFromString(str, v)
}

func (jsonInterator) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (jsonInterator) NewEncoder(writer io.Writer, opts ...EncoderOption) Encoder {
	enc := json.NewEncoder(writer)
	for _, opt := range opts {
		opt(enc)
	}
	return enc
}

func (jsonInterator) NewDecoder(reader io.Reader, opts ...DecoderOption) Decoder {
	dec := json.NewDecoder(reader)
	for _, opt := range opts {
		opt(dec)
	}
	return dec
}

func (jsonInterator) Valid(data []byte) bool {
	return json.Valid(data)
}
