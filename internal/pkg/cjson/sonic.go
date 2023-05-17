//go:build sonic

package cjson

import (
	"io"

	"github.com/bytedance/sonic"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	_jsonAPI            Interface = &jsonSonic{}
	MarshalToString               = _jsonAPI.MarshalToString
	Marshal                       = _jsonAPI.Marshal
	MarshalIndent                 = _jsonAPI.MarshalIndent
	UnmarshalFromString           = _jsonAPI.UnmarshalFromString
	Unmarshal                     = _jsonAPI.Unmarshal
	NewEncoder                    = _jsonAPI.NewEncoder
	NewDecoder                    = _jsonAPI.NewDecoder
	Valid                         = _jsonAPI.Valid

	json = sonic.ConfigStd
)

type jsonSonic struct{}

func (jsonSonic) MarshalToString(v any) (string, error) {
	return json.MarshalToString(v)
}

func (jsonSonic) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonSonic) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (jsonSonic) UnmarshalFromString(str string, v any) error {
	return json.UnmarshalFromString(str, v)
}

func (jsonSonic) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (jsonSonic) NewEncoder(writer io.Writer, opts ...EncoderOption) Encoder {
	enc := json.NewEncoder(writer)
	setEncoderOptions(enc, opts...)
	return enc
}

func (jsonSonic) NewDecoder(reader io.Reader, opts ...DecoderOption) Decoder {
	dec := json.NewDecoder(reader)
	setDecoderOptions(dec, opts...)
	return dec
}

func (jsonSonic) Valid(data []byte) bool {
	return json.Valid(data)
}
