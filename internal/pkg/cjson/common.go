package cjson

import (
	"io"
)

type Interface interface {
	MarshalToString(v any) (string, error)
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
	UnmarshalFromString(str string, v any) error
	Unmarshal(data []byte, v any) error
	NewEncoder(writer io.Writer, opts ...EncoderOption) Encoder
	NewDecoder(reader io.Reader, opts ...DecoderOption) Decoder
	Valid(data []byte) bool
}

type Encoder interface {
	Encode(any) error
	SetEscapeHTML(bool)
}

type EncoderOption func(enc Encoder)

func setEncoderOptions(enc Encoder, opts ...EncoderOption) {
	if enc == nil {
		return
	}
	for _, opt := range opts {
		opt(enc)
	}
}

func SetEscapeHTML(on bool) EncoderOption {
	return func(enc Encoder) {
		enc.SetEscapeHTML(on)
	}
}

type Decoder interface {
	Decode(v any) error
	UseNumber()
}

type DecoderOption func(dec Decoder)

func setDecoderOptions(dec Decoder, opts ...DecoderOption) {
	if dec == nil {
		return
	}
	for _, opt := range opts {
		opt(dec)
	}
}

func UseNumber() DecoderOption {
	return func(dec Decoder) {
		dec.UseNumber()
	}
}
