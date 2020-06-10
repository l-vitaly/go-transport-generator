// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package httpserver

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	_v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver(in *jlexer.Lexer, out *putDetailsTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver(out *jwriter.Writer, in putDetailsTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v putDetailsTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v putDetailsTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *putDetailsTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *putDetailsTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver1(in *jlexer.Lexer, out *putDetailsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Cool":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &out.Cool)
		case "Nothing":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in, &out.Nothing)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver1(out *jwriter.Writer, in putDetailsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Cool\":"
		out.RawString(prefix[1:])
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, in.Cool)
	}
	{
		const prefix string = ",\"Nothing\":"
		out.RawString(prefix)
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out, in.Nothing)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v putDetailsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v putDetailsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *putDetailsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *putDetailsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver1(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in *jlexer.Lexer, out *_v1.Namespace) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out *jwriter.Writer, in _v1.Namespace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in *jlexer.Lexer, out *_v1.Detail) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "namespace":
			out.Namespace = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out *jwriter.Writer, in _v1.Detail) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix[1:])
		out.String(string(in.Namespace))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver2(in *jlexer.Lexer, out *putDetailsRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ThePretty":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &out.Pretty)
		case "Yang":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in, &out.Yang)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver2(out *jwriter.Writer, in putDetailsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ThePretty\":"
		out.RawString(prefix[1:])
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, in.Pretty)
	}
	{
		const prefix string = ",\"Yang\":"
		out.RawString(prefix)
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out, in.Yang)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v putDetailsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v putDetailsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *putDetailsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *putDetailsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver2(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver3(in *jlexer.Lexer, out *getWarehousesTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver3(out *jwriter.Writer, in getWarehousesTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getWarehousesTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getWarehousesTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getWarehousesTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getWarehousesTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver3(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver4(in *jlexer.Lexer, out *getWarehousesResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Pets":
			if in.IsNull() {
				in.Skip()
				out.Pets = nil
			} else {
				in.Delim('[')
				if out.Pets == nil {
					if !in.IsDelim(']') {
						out.Pets = make([]_v1.Detail, 0, 2)
					} else {
						out.Pets = []_v1.Detail{}
					}
				} else {
					out.Pets = (out.Pets)[:0]
				}
				for !in.IsDelim(']') {
					var v1 _v1.Detail
					easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &v1)
					out.Pets = append(out.Pets, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver4(out *jwriter.Writer, in getWarehousesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Pets\":"
		out.RawString(prefix[1:])
		if in.Pets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Pets {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getWarehousesResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getWarehousesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getWarehousesResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getWarehousesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver4(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver5(in *jlexer.Lexer, out *getSomeElseDataUtf8Transport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver5(out *jwriter.Writer, in getSomeElseDataUtf8Transport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getSomeElseDataUtf8Transport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getSomeElseDataUtf8Transport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getSomeElseDataUtf8Transport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getSomeElseDataUtf8Transport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver5(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver6(in *jlexer.Lexer, out *getSomeElseDataUtf8Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cool":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &out.Cool)
		case "id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(string)
				}
				*out.Id = string(in.String())
			}
		case "TheNothing":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in, &out.Nothing)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver6(out *jwriter.Writer, in getSomeElseDataUtf8Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cool\":"
		out.RawString(prefix[1:])
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, in.Cool)
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		if in.Id == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Id))
		}
	}
	{
		const prefix string = ",\"TheNothing\":"
		out.RawString(prefix)
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out, in.Nothing)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getSomeElseDataUtf8Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getSomeElseDataUtf8Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getSomeElseDataUtf8Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getSomeElseDataUtf8Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver6(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver7(in *jlexer.Lexer, out *getDetailsTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver7(out *jwriter.Writer, in getDetailsTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getDetailsTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getDetailsTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getDetailsTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getDetailsTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver7(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver8(in *jlexer.Lexer, out *getDetailsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Det":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &out.Det)
		case "Ns":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in, &out.Ns)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver8(out *jwriter.Writer, in getDetailsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Det\":"
		out.RawString(prefix[1:])
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, in.Det)
	}
	{
		const prefix string = ",\"Ns\":"
		out.RawString(prefix)
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out, in.Ns)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getDetailsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getDetailsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getDetailsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getDetailsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver8(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver9(in *jlexer.Lexer, out *getDetailsListEmbedStructTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver9(out *jwriter.Writer, in getDetailsListEmbedStructTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getDetailsListEmbedStructTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getDetailsListEmbedStructTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getDetailsListEmbedStructTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getDetailsListEmbedStructTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver9(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver10(in *jlexer.Lexer, out *getDetailsListEmbedStructResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(getDetailsListEmbedStructResponse, 0, 2)
			} else {
				*out = getDetailsListEmbedStructResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 _v1.Detail
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &v4)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver10(out *jwriter.Writer, in getDetailsListEmbedStructResponse) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, v6)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v getDetailsListEmbedStructResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getDetailsListEmbedStructResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getDetailsListEmbedStructResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getDetailsListEmbedStructResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver10(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver11(in *jlexer.Lexer, out *getDetailsEmbedStructTransport) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver11(out *jwriter.Writer, in getDetailsEmbedStructTransport) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getDetailsEmbedStructTransport) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getDetailsEmbedStructTransport) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getDetailsEmbedStructTransport) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getDetailsEmbedStructTransport) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver11(l, v)
}
func easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver12(in *jlexer.Lexer, out *getDetailsEmbedStructResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "detail":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(in, &out.Detail)
		case "namespace":
			easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(in, &out.Namespace)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver12(out *jwriter.Writer, in getDetailsEmbedStructResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"detail\":"
		out.RawString(prefix[1:])
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV1(out, in.Detail)
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleApiV11(out, in.Namespace)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v getDetailsEmbedStructResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v getDetailsEmbedStructResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC87d08bdEncodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *getDetailsEmbedStructResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *getDetailsEmbedStructResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC87d08bdDecodeGithubComWildberriesRuGoTransportGeneratorExampleServiceHttpserver12(l, v)
}
