//line parser.go.y:2
package gql

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"fmt"
	"strconv"
)

type Token struct {
	token   int
	literal string
}

type Expression interface{}

type SelectExpr struct {
	Field  FieldExpr
	From   *FromExpr
	Where  []ConditionExpr
	Order  []OrderExpr
	Limit  *LimitExpr
	Offset *OffsetExpr
}

type FieldExpr struct {
	Distinct        bool
	DistinctOnField []string
	Field           []string
	Asterisk        bool
}

type FromExpr struct {
	Kind *KindExpr
}

type ConditionExpr interface {
	GetPropertyName() string
	GetValue() ValueExpr
	GetComparator() ComparatorExpr
}

type IsNullConditionExpr struct {
	PropertyName string
}

func (c IsNullConditionExpr) GetPropertyName() string {
	return c.PropertyName
}

func (c IsNullConditionExpr) GetValue() ValueExpr {
	return ValueExpr{}
}

func (c IsNullConditionExpr) GetComparator() ComparatorExpr {
	return OP_IS_NULL
}

type ForwardConditionExpr struct {
	PropertyName string
	Comparator   ComparatorExpr
	Value        ValueExpr
}

func (c ForwardConditionExpr) GetPropertyName() string {
	return c.PropertyName
}

func (c ForwardConditionExpr) GetValue() ValueExpr {
	return c.Value
}

func (c ForwardConditionExpr) GetComparator() ComparatorExpr {
	return c.Comparator
}

type BackwardConditionExpr struct {
	Value        ValueExpr
	Comparator   ComparatorExpr
	PropertyName string
}

func (c BackwardConditionExpr) GetPropertyName() string {
	return c.PropertyName
}

func (c BackwardConditionExpr) GetValue() ValueExpr {
	return c.Value
}

func (c BackwardConditionExpr) GetComparator() ComparatorExpr {
	return c.Comparator
}

type OrderExpr struct {
	PropertyName string
	Sort         SortType
}

type LimitExpr struct {
	Cursor string
	Number int
}

type OffsetExpr struct {
	Cursor string
	Number int
}

type ResultPositionExpr struct {
	Number      int
	BindingSite string
}

type ValueExpr struct {
	Type ValueType
	V    interface{}
}

type KeyLiteralExpr struct {
	Project   string
	Namespace string
	KeyPath   []KeyPathElementExpr
}

type BlobLiteralExpr struct {
	Blob string
}

type DatetimeLiteralExpr struct {
	Datetime string
}

type KeyPathElementExpr struct {
	Kind string
	ID   int64
	Name string
}

type KindExpr struct {
	Name string
}

type PropertyNameExpr struct {
	Name string
}

type ComparatorExpr int

func (c ComparatorExpr) String() string {
	switch c {
	case OP_IS_NULL:
		return "IS NULL"
	case OP_CONTAINS:
		return "CONTAINS"
	case OP_HAS_ANCESTOR:
		return "HAS ANCESTOR"
	case OP_IN:
		return "IN"
	case OP_HAS_DESCENDANT:
		return "HAS DESCENDANT"
	case OP_EQUALS:
		return "="
	case OP_LESS:
		return "<"
	case OP_LESS_EQUALS:
		return "<="
	case OP_GREATER:
		return ">"
	case OP_GREATER_EQUALS:
		return ">="
	}
	return ""
}

const (
	OP_IS_NULL ComparatorExpr = 1 << iota
	OP_CONTAINS
	OP_HAS_ANCESTOR
	OP_IN
	OP_HAS_DESCENDANT
	OP_EQUALS         // =
	OP_LESS           // <
	OP_LESS_EQUALS    // <=
	OP_GREATER        // >
	OP_GREATER_EQUALS // >=
)

type ValueType int

const (
	TYPE_BINDING_SITE ValueType = 1 << iota
	TYPE_KEY
	TYPE_BLOB
	TYPE_DATETIME
	TYPE_STRING
	TYPE_INTEGER
	TYPE_DOUBLE
	TYPE_BOOL
	TYPE_NULL
)

type SortType int

const (
	SORT_NONE SortType = iota
	SORT_ASC
	SORT_DESC
)

//line parser.go.y:212
type yySymType struct {
	yys   int
	token Token
	expr  Expression
}

const ILLEGAL = 57346
const EOF = 57347
const WS = 57348
const NAME = 57349
const BINDING_SITE = 57350
const STRING = 57351
const INTEGER = 57352
const DOUBLE = 57353
const ASTERISK = 57354
const PLUS = 57355
const COMMA = 57356
const EQUAL = 57357
const LEFT_BRACKETS = 57358
const RIGHT_BRACKETS = 57359
const LEFT_ROUND = 57360
const RIGHT_ROUND = 57361
const SELECT = 57362
const DISTINCT = 57363
const ON = 57364
const FROM = 57365
const WHERE = 57366
const ASC = 57367
const DESC = 57368
const ORDER = 57369
const BY = 57370
const LIMIT = 57371
const FIRST = 57372
const OFFSET = 57373
const AND = 57374
const IS = 57375
const NULL = 57376
const CONTAINS = 57377
const HAS = 57378
const ANCESTOR = 57379
const DESCENDANT = 57380
const IN = 57381
const KEY = 57382
const PROJECT = 57383
const NAMESPACE = 57384
const BLOB = 57385
const DATETIME = 57386
const TRUE = 57387
const FALSE = 57388

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ILLEGAL",
	"EOF",
	"WS",
	"NAME",
	"BINDING_SITE",
	"STRING",
	"INTEGER",
	"DOUBLE",
	"ASTERISK",
	"PLUS",
	"COMMA",
	"EQUAL",
	"LEFT_BRACKETS",
	"RIGHT_BRACKETS",
	"LEFT_ROUND",
	"RIGHT_ROUND",
	"SELECT",
	"DISTINCT",
	"ON",
	"FROM",
	"WHERE",
	"ASC",
	"DESC",
	"ORDER",
	"BY",
	"LIMIT",
	"FIRST",
	"OFFSET",
	"AND",
	"IS",
	"NULL",
	"CONTAINS",
	"HAS",
	"ANCESTOR",
	"DESCENDANT",
	"IN",
	"KEY",
	"PROJECT",
	"NAMESPACE",
	"BLOB",
	"DATETIME",
	"TRUE",
	"FALSE",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:702
type Lexer struct {
	Scanner    *Scanner
	Result     Expression
	parserErr  error
	scannerErr error
}

func (l *Lexer) Lex(lval *yySymType) int {
	token, literal := l.Scanner.Scan()

	if token == EOF {
		return 0

	} else if token == ILLEGAL {
		l.scannerErr = fmt.Errorf("invalid token: %v", literal)
		return 0
	}

	lval.token = Token{token: int(token), literal: literal}
	return int(token)
}

func (l *Lexer) Error(e string) {
	if l.scannerErr == nil {
		l.parserErr = fmt.Errorf("%v: %v\n", e, l.Scanner.Consumed())
	}
}

func (l *Lexer) Parse() error {
	yyParse(l)

	if l.parserErr != nil {
		return l.parserErr
	} else {
		return l.scannerErr
	}
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 66
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 136

var yyAct = [...]int{

	96, 61, 84, 8, 6, 9, 27, 29, 30, 31,
	88, 24, 13, 76, 74, 26, 19, 70, 68, 42,
	25, 27, 29, 30, 31, 38, 60, 48, 49, 50,
	40, 41, 34, 85, 86, 16, 22, 45, 35, 11,
	3, 36, 37, 32, 33, 66, 25, 34, 54, 114,
	63, 53, 64, 35, 67, 73, 36, 37, 32, 33,
	69, 9, 81, 80, 52, 48, 49, 50, 9, 104,
	12, 112, 62, 5, 103, 58, 14, 107, 91, 90,
	98, 89, 7, 43, 93, 46, 47, 94, 82, 57,
	56, 55, 20, 72, 100, 71, 115, 102, 113, 105,
	101, 12, 83, 108, 92, 109, 9, 97, 18, 111,
	110, 79, 63, 95, 64, 106, 99, 78, 77, 17,
	9, 87, 75, 28, 51, 44, 23, 59, 65, 21,
	39, 15, 10, 4, 2, 1,
}
var yyPact = [...]int{

	20, -1000, -1000, 61, 16, -1000, 87, 54, -1000, -1000,
	11, 101, 113, 87, 74, 9, -2, -1000, -1000, -1000,
	113, 1, 3, -13, -1000, 50, 12, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 73, 72, 71, 56, -5,
	42, 113, -2, -16, 13, -1000, -1000, -20, -1000, 80,
	78, 113, -1000, -1000, -24, -28, 109, 108, 99, -1000,
	104, -1000, 70, -1000, -1000, 88, 8, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -32, 63, 60, 59, -1000,
	87, 91, 104, 113, -1000, -1000, -1000, 101, 62, 107,
	-1000, -1000, 104, 86, 8, 55, -1000, 85, 106, 58,
	-1000, 104, -1000, -1000, 101, 100, 52, 84, 30, -1000,
	-1000, -1000, 82, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 135, 134, 133, 4, 132, 131, 130, 129, 128,
	2, 127, 126, 11, 125, 124, 37, 1, 123, 122,
	121, 113, 0, 107, 3, 15,
}
var yyR1 = [...]int{

	0, 1, 2, 3, 3, 3, 3, 3, 4, 4,
	5, 5, 6, 6, 12, 12, 13, 13, 13, 14,
	14, 14, 15, 15, 15, 16, 16, 16, 16, 16,
	8, 8, 9, 9, 10, 10, 10, 7, 7, 7,
	11, 11, 11, 17, 17, 25, 25, 25, 25, 25,
	25, 25, 25, 18, 18, 18, 19, 19, 20, 20,
	21, 21, 22, 22, 23, 24,
}
var yyR2 = [...]int{

	0, 1, 7, 1, 1, 2, 6, 6, 1, 3,
	0, 2, 0, 2, 1, 3, 3, 3, 3, 1,
	1, 2, 1, 1, 2, 1, 1, 2, 1, 2,
	0, 3, 2, 4, 0, 1, 1, 0, 2, 7,
	0, 2, 4, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 6, 4, 4, 0, 5, 0, 5,
	1, 3, 3, 3, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, 20, -3, 12, -4, 21, -24, 7,
	-5, 23, 14, -4, 22, -6, 24, -23, 7, -24,
	18, -8, 27, -12, -13, -24, -25, 8, -18, 9,
	10, 11, 45, 46, 34, 40, 43, 44, -4, -7,
	29, 28, 32, 33, -14, -16, 35, 36, 15, 16,
	17, -15, -16, 39, 36, 18, 18, 18, 19, -11,
	31, -17, 30, 8, 10, -9, -24, -13, 34, -25,
	37, 15, 15, -24, 38, -19, 41, 9, 9, 12,
	-4, -17, 18, 14, -10, 25, 26, -20, 42, 18,
	19, 19, 13, -17, -24, -21, -22, -23, 18, 9,
	-17, 14, -10, 19, 14, 14, 9, 19, -17, -22,
	10, 9, 19, 14, 19, 14,
}
var yyDef = [...]int{

	0, -2, 1, 0, 10, 3, 4, 0, 8, 65,
	12, 0, 0, 5, 0, 30, 0, 11, 64, 9,
	0, 37, 0, 13, 14, 0, 0, 45, 46, 47,
	48, 49, 50, 51, 52, 0, 0, 0, 0, 40,
	0, 0, 0, 0, 0, 19, 20, 0, 25, 26,
	28, 0, 22, 23, 0, 56, 0, 0, 0, 2,
	0, 38, 0, 43, 44, 31, 34, 15, 16, 17,
	21, 27, 29, 18, 24, 58, 0, 0, 0, 6,
	7, 41, 0, 0, 32, 35, 36, 0, 0, 0,
	54, 55, 0, 0, 34, 0, 60, 0, 0, 0,
	42, 0, 33, 53, 0, 0, 0, 0, 0, 61,
	62, 63, 0, 57, 39, 59,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:300
		{
			yyVAL.expr = yyDollar[1].expr
			yylex.(*Lexer).Result = yyVAL.expr
		}
	case 2:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:307
		{
			fieldExpr := yyDollar[2].expr
			fromExpr := yyDollar[3].expr
			whereExpr := yyDollar[4].expr
			orderExpr := yyDollar[5].expr
			limitExpr := yyDollar[6].expr
			offsetExpr := yyDollar[7].expr

			var from *FromExpr
			if fromExpr != nil {
				from = fromExpr.(*FromExpr)
			}

			var limit *LimitExpr
			if limitExpr != nil {
				limit = limitExpr.(*LimitExpr)
			}

			var offset *OffsetExpr
			if offsetExpr != nil {
				offset = offsetExpr.(*OffsetExpr)
			}

			yyVAL.expr = SelectExpr{
				Field:  fieldExpr.(FieldExpr),
				From:   from,
				Where:  whereExpr.([]ConditionExpr),
				Order:  orderExpr.([]OrderExpr),
				Limit:  limit,
				Offset: offset,
			}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.expr = FieldExpr{Asterisk: true}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr = FieldExpr{Field: yyDollar[1].expr.([]string)}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:350
		{
			yyVAL.expr = FieldExpr{
				Distinct: true,
				Field:    yyDollar[2].expr.([]string),
			}
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:357
		{
			yyVAL.expr = FieldExpr{
				DistinctOnField: yyDollar[4].expr.([]string),
				Asterisk:        true,
			}
		}
	case 7:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr = FieldExpr{
				DistinctOnField: yyDollar[4].expr.([]string),
				Field:           yyDollar[6].expr.([]string),
			}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr = []string{yyDollar[1].expr.(string)}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:377
		{
			yyVAL.expr = append(yyDollar[1].expr.([]string), yyDollar[3].expr.(string))
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr = nil
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:387
		{
			kind := yyDollar[2].expr.(KindExpr)
			yyVAL.expr = &FromExpr{Kind: &kind}
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.expr = make([]ConditionExpr, 0)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:398
		{
			yyVAL.expr = yyDollar[2].expr.([]ConditionExpr)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:404
		{
			yyVAL.expr = []ConditionExpr{yyDollar[1].expr.(ConditionExpr)}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:408
		{
			yyVAL.expr = append(yyDollar[1].expr.([]ConditionExpr), yyDollar[3].expr.(ConditionExpr))
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:414
		{
			yyVAL.expr = IsNullConditionExpr{
				PropertyName: yyDollar[1].expr.(string),
			}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:420
		{
			yyVAL.expr = ForwardConditionExpr{
				PropertyName: yyDollar[1].expr.(string),
				Comparator:   yyDollar[2].expr.(ComparatorExpr),
				Value:        yyDollar[3].expr.(ValueExpr),
			}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:428
		{
			yyVAL.expr = BackwardConditionExpr{
				Value:        yyDollar[1].expr.(ValueExpr),
				Comparator:   yyDollar[2].expr.(ComparatorExpr),
				PropertyName: yyDollar[3].expr.(string),
			}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:438
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:442
		{
			yyVAL.expr = OP_CONTAINS
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:446
		{
			yyVAL.expr = OP_HAS_ANCESTOR
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:452
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:456
		{
			yyVAL.expr = OP_IN
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:460
		{
			yyVAL.expr = OP_HAS_DESCENDANT
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:466
		{
			yyVAL.expr = OP_EQUALS
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:470
		{
			yyVAL.expr = OP_LESS
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:474
		{
			yyVAL.expr = OP_LESS_EQUALS
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:478
		{
			yyVAL.expr = OP_GREATER
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = OP_GREATER_EQUALS
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:488
		{
			yyVAL.expr = []OrderExpr{}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:492
		{
			yyVAL.expr = yyDollar[3].expr.([]OrderExpr)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:498
		{
			yyVAL.expr = []OrderExpr{
				OrderExpr{PropertyName: yyDollar[1].expr.(string), Sort: yyDollar[2].expr.(SortType)},
			}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:504
		{
			o := OrderExpr{PropertyName: yyDollar[3].expr.(string), Sort: yyDollar[4].expr.(SortType)}
			yyVAL.expr = append(yyDollar[1].expr.([]OrderExpr), o)
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:511
		{
			yyVAL.expr = SORT_NONE
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:515
		{
			yyVAL.expr = SORT_ASC
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:519
		{
			yyVAL.expr = SORT_DESC
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:525
		{
			yyVAL.expr = nil
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:529
		{
			yyVAL.expr = &LimitExpr{
				Cursor: yyDollar[2].expr.(ResultPositionExpr).BindingSite,
				Number: yyDollar[2].expr.(ResultPositionExpr).Number,
			}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:536
		{
			yyVAL.expr = &LimitExpr{
				Cursor: yyDollar[4].expr.(ResultPositionExpr).BindingSite,
				Number: yyDollar[6].expr.(ResultPositionExpr).Number,
			}
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:545
		{
			yyVAL.expr = nil
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:549
		{
			if yyDollar[2].expr.(ResultPositionExpr).BindingSite != "" {
				yyVAL.expr = &OffsetExpr{Cursor: yyDollar[2].expr.(ResultPositionExpr).BindingSite}
			} else {
				yyVAL.expr = &OffsetExpr{Number: yyDollar[2].expr.(ResultPositionExpr).Number}
			}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.expr = &OffsetExpr{
				Cursor: yyDollar[2].expr.(ResultPositionExpr).BindingSite,
				Number: yyDollar[4].expr.(ResultPositionExpr).Number,
			}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:566
		{
			yyVAL.expr = ResultPositionExpr{BindingSite: yyDollar[1].token.literal}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:570
		{
			number, err := strconv.Atoi(yyDollar[1].token.literal)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to integer", yyDollar[1].token.literal))
			}
			yyVAL.expr = ResultPositionExpr{Number: number}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:580
		{
			yyVAL.expr = ValueExpr{Type: TYPE_BINDING_SITE, V: yyDollar[1].token.literal}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:584
		{
			switch yyDollar[1].expr.(type) {
			case KeyLiteralExpr:
				yyVAL.expr = ValueExpr{Type: TYPE_KEY, V: yyDollar[1].expr}
			case BlobLiteralExpr:
				yyVAL.expr = ValueExpr{Type: TYPE_BLOB, V: yyDollar[1].expr}
			case DatetimeLiteralExpr:
				yyVAL.expr = ValueExpr{Type: TYPE_DATETIME, V: yyDollar[1].expr}
			default:
				panic(fmt.Sprintf("unkown synthetic_literal:%v", yyDollar[1].expr))
			}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:597
		{
			yyVAL.expr = ValueExpr{Type: TYPE_STRING, V: yyDollar[1].token.literal}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:601
		{
			number, err := strconv.ParseInt(yyDollar[1].token.literal, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to integer", yyDollar[1].token.literal))
			}
			yyVAL.expr = ValueExpr{Type: TYPE_INTEGER, V: number}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:609
		{
			double, err := strconv.ParseFloat(yyDollar[1].token.literal, 64)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to double", yyDollar[1].token.literal))
			}
			yyVAL.expr = ValueExpr{Type: TYPE_DOUBLE, V: double}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:617
		{
			yyVAL.expr = ValueExpr{Type: TYPE_BOOL, V: true}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:621
		{
			yyVAL.expr = ValueExpr{Type: TYPE_BOOL, V: false}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:625
		{
			yyVAL.expr = ValueExpr{Type: TYPE_NULL, V: nil}
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:631
		{
			yyVAL.expr = KeyLiteralExpr{
				Project:   yyDollar[3].expr.(string),
				Namespace: yyDollar[4].expr.(string),
				KeyPath:   yyDollar[5].expr.([]KeyPathElementExpr),
			}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:639
		{
			yyVAL.expr = BlobLiteralExpr{Blob: yyDollar[3].token.literal}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:643
		{
			yyVAL.expr = DatetimeLiteralExpr{Datetime: yyDollar[3].token.literal}
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:649
		{
			yyVAL.expr = nil
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:653
		{
			yyVAL.expr = yyDollar[3].token.literal
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:659
		{
			yyVAL.expr = nil
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:663
		{
			yyVAL.expr = yyDollar[3].token.literal
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:669
		{
			yyVAL.expr = []KeyPathElementExpr{yyDollar[1].expr.(KeyPathElementExpr)}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:673
		{
			yyVAL.expr = append(yyDollar[1].expr.([]KeyPathElementExpr), yyDollar[3].expr.(KeyPathElementExpr))
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:679
		{
			number, err := strconv.ParseInt(yyDollar[3].token.literal, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to integer", yyDollar[3].token.literal))
			}
			yyVAL.expr = KeyPathElementExpr{Kind: yyDollar[1].expr.(KindExpr).Name, ID: number}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:687
		{
			yyVAL.expr = KeyPathElementExpr{Kind: yyDollar[1].expr.(KindExpr).Name, Name: yyDollar[3].token.literal}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:693
		{
			yyVAL.expr = KindExpr{Name: yyDollar[1].token.literal}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:699
		{
			yyVAL.expr = yyDollar[1].token.literal
		}
	}
	goto yystack /* stack new state and value */
}
