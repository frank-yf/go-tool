//go:build !jsoniter && !sonic

package cjson

import (
	"encoding/json"
	"io"

	"github.com/frank-yf/go-tool/internal/pkg/tool"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	_json               Interface = &std{}
	MarshalToString               = _json.MarshalToString
	Marshal                       = _json.Marshal
	MarshalIndent                 = _json.MarshalIndent
	UnmarshalFromString           = _json.UnmarshalFromString
	Unmarshal                     = _json.Unmarshal
	NewEncoder                    = _json.NewEncoder
	NewDecoder                    = _json.NewDecoder
	Valid                         = _json.Valid
)

type std struct{}

func (s std) MarshalToString(v any) (string, error) {
	bs, err := s.Marshal(v)
	if err != nil {
		return "", err
	}
	return tool.BytesToString(bs), nil
}

func (std) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (std) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (s std) UnmarshalFromString(str string, v any) error {
	bs := tool.StringToBytes(str)
	return s.Unmarshal(bs, v)
}

func (std) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (std) NewEncoder(writer io.Writer, opts ...EncoderOption) Encoder {
	enc := json.NewEncoder(writer)
	for _, opt := range opts {
		opt(enc)
	}
	return enc
}

func (std) NewDecoder(reader io.Reader, opts ...DecoderOption) Decoder {
	dec := json.NewDecoder(reader)
	for _, opt := range opts {
		opt(dec)
	}
	return dec
}

func (s std) Valid(data []byte) bool {
	var v any
	err := s.Unmarshal(data, &v)
	return err == nil
}
