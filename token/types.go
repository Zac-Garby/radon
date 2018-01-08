package token

// Type is the type of a token
type Type string

// All the possible types of tokens.
const (
	EOF = "EOF"

	Illegal = "illegal"
	Number  = "number"
	String  = "string"
	ID      = "identifier"

	Plus           = "plus"
	Minus          = "minus"
	Star           = "star"
	Exp            = "exponent"
	Slash          = "slash"
	FloorDiv       = "floor-div"
	Mod            = "modulo"
	LeftParen      = "left-paren"
	RightParen     = "right-paren"
	LessThan       = "less-than"
	GreaterThan    = "greater-than"
	LessThanEq     = "less-than-or-equal"
	GreaterThanEq  = "greater-than-or-equal"
	LeftBrace      = "left-brace"
	RightBrace     = "right-brace"
	LeftSquare     = "left-square"
	RightSquare    = "right-square"
	Semi           = "semi"
	Equal          = "equal"
	NotEqual       = "not-equal"
	Or             = "or"
	And            = "and"
	BitOr          = "bitwise-or"
	BitAnd         = "bitwise-and"
	Assign         = "assign"
	Declare        = "declare"
	Comma          = "comma"
	RightArrow     = "right-arrow"
	LambdaArrow    = "lambda-arrow"
	Colon          = "colon"
	Dot            = "dot"
	Bang           = "bang"
	PlusEquals     = "assign-plus"
	MinusEquals    = "assign-minus"
	StarEquals     = "assign-star"
	ExpEquals      = "assign-exponent"
	SlashEquals    = "assign-slash"
	FloorDivEquals = "assign-floor-div"
	ModEquals      = "assign-modulo"
	OrEquals       = "assign-or"
	AndEquals      = "assign-and"
	BitOrEquals    = "assign-bitwise-or"
	BitAndEquals   = "assign-bitwise-and"

	Return = "return"
	True   = "true"
	False  = "false"
	Nil    = "nil"
	If     = "if"
	Then   = "then"
	Else   = "else"
	While  = "while"
	For    = "for"
	Loop   = "loop"
	Do     = "do"
	Next   = "next"
	Break  = "break"
	Match  = "match"
	Model  = "model"
	Map    = "map"
	Where  = "where"
	Import = "import"
)
