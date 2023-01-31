package token

const NoPos = 0

// Token 代表单个已解析的标识
type Token struct {
	Kind     Kind
	Value    interface{}
	Position int
}

var keywords = map[string]Kind{
	"true":  BoolLiteral,
	"false": BoolLiteral,
}

func LookupOperator(op string) Kind {
	if kind, exist := operatorToKind[op]; exist {
		return kind
	}
	return Illegal
}

// Lookup - maps an identifier to its keyword token or Identifier (if not a keyword).
// 将标识符映射到其`关键字 token` 或 标识符(如果不是关键字)。
func Lookup(ident string) Kind {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}

	return Identifier
}
