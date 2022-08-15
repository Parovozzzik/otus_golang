package hw10programoptimization

import (
	jlexer "github.com/mailru/easyjson/jlexer"
)

var _ *jlexer.Lexer

func easyjsonE3ab7953DecodeGithubComFixmeMyFriendHw10ProgramOptimization(in *jlexer.Lexer, out *Email) {
	if in.IsNull() {
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Email":
			*out = Email(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}

func (v *Email) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3ab7953DecodeGithubComFixmeMyFriendHw10ProgramOptimization(l, v)
}
