// Code generated from Age.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Age

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 82, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 23, 10, 2, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 46, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 52, 10, 7, 12, 7, 14, 7, 55, 11, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 61, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 7, 9, 71, 10, 9, 12, 9, 14, 9, 74, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 80, 10, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2, 2, 85,
	2, 22, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 30, 3, 2, 2,
	2, 10, 45, 3, 2, 2, 2, 12, 60, 3, 2, 2, 2, 14, 62, 3, 2, 2, 2, 16, 79,
	3, 2, 2, 2, 18, 23, 5, 10, 6, 2, 19, 23, 5, 4, 3, 2, 20, 23, 5, 6, 4, 2,
	21, 23, 5, 8, 5, 2, 22, 18, 3, 2, 2, 2, 22, 19, 3, 2, 2, 2, 22, 20, 3,
	2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 3, 3, 2, 2, 2, 24, 25, 5, 12, 7, 2, 25,
	26, 7, 9, 2, 2, 26, 5, 3, 2, 2, 2, 27, 28, 5, 12, 7, 2, 28, 29, 7, 10,
	2, 2, 29, 7, 3, 2, 2, 2, 30, 31, 7, 3, 2, 2, 31, 32, 5, 4, 3, 2, 32, 33,
	7, 4, 2, 2, 33, 34, 5, 6, 4, 2, 34, 35, 7, 4, 2, 2, 35, 36, 5, 4, 3, 2,
	36, 37, 7, 5, 2, 2, 37, 38, 7, 11, 2, 2, 38, 9, 3, 2, 2, 2, 39, 46, 7,
	12, 2, 2, 40, 46, 7, 15, 2, 2, 41, 46, 7, 13, 2, 2, 42, 46, 7, 14, 2, 2,
	43, 46, 5, 12, 7, 2, 44, 46, 5, 16, 9, 2, 45, 39, 3, 2, 2, 2, 45, 40, 3,
	2, 2, 2, 45, 41, 3, 2, 2, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45,
	44, 3, 2, 2, 2, 46, 11, 3, 2, 2, 2, 47, 48, 7, 6, 2, 2, 48, 53, 5, 14,
	8, 2, 49, 50, 7, 4, 2, 2, 50, 52, 5, 14, 8, 2, 51, 49, 3, 2, 2, 2, 52,
	55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2,
	2, 55, 53, 3, 2, 2, 2, 56, 57, 7, 7, 2, 2, 57, 61, 3, 2, 2, 2, 58, 59,
	7, 6, 2, 2, 59, 61, 7, 7, 2, 2, 60, 47, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2,
	61, 13, 3, 2, 2, 2, 62, 63, 7, 12, 2, 2, 63, 64, 7, 8, 2, 2, 64, 65, 5,
	10, 6, 2, 65, 15, 3, 2, 2, 2, 66, 67, 7, 3, 2, 2, 67, 72, 5, 10, 6, 2,
	68, 69, 7, 4, 2, 2, 69, 71, 5, 10, 6, 2, 70, 68, 3, 2, 2, 2, 71, 74, 3,
	2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74,
	72, 3, 2, 2, 2, 75, 76, 7, 5, 2, 2, 76, 80, 3, 2, 2, 2, 77, 78, 7, 3, 2,
	2, 78, 80, 7, 5, 2, 2, 79, 66, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 17,
	3, 2, 2, 2, 8, 22, 45, 53, 60, 72, 79,
}
var literalNames = []string{
	"", "'['", "','", "']'", "'{'", "'}'", "':'", "'::vertex'", "'::edge'",
	"'::path'", "", "", "'null'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "KW_VERTEX", "KW_EDGE", "KW_PATH", "STRING",
	"BOOL", "NULL", "NUMBER", "WS",
}

var ruleNames = []string{
	"ageout", "vertex", "edge", "path", "value", "properties", "pair", "arr",
}

type AgeParser struct {
	*antlr.BaseParser
}

// NewAgeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *AgeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAgeParser(input antlr.TokenStream) *AgeParser {
	this := new(AgeParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Age.g4"

	return this
}

// AgeParser tokens.
const (
	AgeParserEOF       = antlr.TokenEOF
	AgeParserT__0      = 1
	AgeParserT__1      = 2
	AgeParserT__2      = 3
	AgeParserT__3      = 4
	AgeParserT__4      = 5
	AgeParserT__5      = 6
	AgeParserKW_VERTEX = 7
	AgeParserKW_EDGE   = 8
	AgeParserKW_PATH   = 9
	AgeParserSTRING    = 10
	AgeParserBOOL      = 11
	AgeParserNULL      = 12
	AgeParserNUMBER    = 13
	AgeParserWS        = 14
)

// AgeParser rules.
const (
	AgeParserRULE_ageout     = 0
	AgeParserRULE_vertex     = 1
	AgeParserRULE_edge       = 2
	AgeParserRULE_path       = 3
	AgeParserRULE_value      = 4
	AgeParserRULE_properties = 5
	AgeParserRULE_pair       = 6
	AgeParserRULE_arr        = 7
)

// IAgeoutContext is an interface to support dynamic dispatch.
type IAgeoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgeoutContext differentiates from other interfaces.
	IsAgeoutContext()
}

type AgeoutContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgeoutContext() *AgeoutContext {
	var p = new(AgeoutContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_ageout
	return p
}

func (*AgeoutContext) IsAgeoutContext() {}

func NewAgeoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgeoutContext {
	var p = new(AgeoutContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_ageout

	return p
}

func (s *AgeoutContext) GetParser() antlr.Parser { return s.parser }

func (s *AgeoutContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AgeoutContext) Vertex() IVertexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertexContext)
}

func (s *AgeoutContext) Edge() IEdgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdgeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdgeContext)
}

func (s *AgeoutContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *AgeoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgeoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgeoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterAgeout(s)
	}
}

func (s *AgeoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitAgeout(s)
	}
}

func (s *AgeoutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitAgeout(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Ageout() (localctx IAgeoutContext) {
	localctx = NewAgeoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgeParserRULE_ageout)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Vertex()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(18)
			p.Edge()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(19)
			p.Path()
		}

	}

	return localctx
}

// IVertexContext is an interface to support dynamic dispatch.
type IVertexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVertexContext differentiates from other interfaces.
	IsVertexContext()
}

type VertexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertexContext() *VertexContext {
	var p = new(VertexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_vertex
	return p
}

func (*VertexContext) IsVertexContext() {}

func NewVertexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertexContext {
	var p = new(VertexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_vertex

	return p
}

func (s *VertexContext) GetParser() antlr.Parser { return s.parser }

func (s *VertexContext) Properties() IPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *VertexContext) KW_VERTEX() antlr.TerminalNode {
	return s.GetToken(AgeParserKW_VERTEX, 0)
}

func (s *VertexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VertexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterVertex(s)
	}
}

func (s *VertexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitVertex(s)
	}
}

func (s *VertexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitVertex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Vertex() (localctx IVertexContext) {
	localctx = NewVertexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgeParserRULE_vertex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Properties()
	}
	{
		p.SetState(23)
		p.Match(AgeParserKW_VERTEX)
	}

	return localctx
}

// IEdgeContext is an interface to support dynamic dispatch.
type IEdgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeContext differentiates from other interfaces.
	IsEdgeContext()
}

type EdgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeContext() *EdgeContext {
	var p = new(EdgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_edge
	return p
}

func (*EdgeContext) IsEdgeContext() {}

func NewEdgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeContext {
	var p = new(EdgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_edge

	return p
}

func (s *EdgeContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeContext) Properties() IPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *EdgeContext) KW_EDGE() antlr.TerminalNode {
	return s.GetToken(AgeParserKW_EDGE, 0)
}

func (s *EdgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterEdge(s)
	}
}

func (s *EdgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitEdge(s)
	}
}

func (s *EdgeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitEdge(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Edge() (localctx IEdgeContext) {
	localctx = NewEdgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AgeParserRULE_edge)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Properties()
	}
	{
		p.SetState(26)
		p.Match(AgeParserKW_EDGE)
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllVertex() []IVertexContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVertexContext)(nil)).Elem())
	var tst = make([]IVertexContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVertexContext)
		}
	}

	return tst
}

func (s *PathContext) Vertex(i int) IVertexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertexContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVertexContext)
}

func (s *PathContext) Edge() IEdgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdgeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdgeContext)
}

func (s *PathContext) KW_PATH() antlr.TerminalNode {
	return s.GetToken(AgeParserKW_PATH, 0)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AgeParserRULE_path)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(AgeParserT__0)
	}
	{
		p.SetState(29)
		p.Vertex()
	}
	{
		p.SetState(30)
		p.Match(AgeParserT__1)
	}
	{
		p.SetState(31)
		p.Edge()
	}
	{
		p.SetState(32)
		p.Match(AgeParserT__1)
	}
	{
		p.SetState(33)
		p.Vertex()
	}
	{
		p.SetState(34)
		p.Match(AgeParserT__2)
	}
	{
		p.SetState(35)
		p.Match(AgeParserKW_PATH)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgeParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AgeParserNUMBER, 0)
}

func (s *ValueContext) BOOL() antlr.TerminalNode {
	return s.GetToken(AgeParserBOOL, 0)
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(AgeParserNULL, 0)
}

func (s *ValueContext) Properties() IPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertiesContext)
}

func (s *ValueContext) Arr() IArrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AgeParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AgeParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(AgeParserSTRING)
		}

	case AgeParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(AgeParserNUMBER)
		}

	case AgeParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Match(AgeParserBOOL)
		}

	case AgeParserNULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Match(AgeParserNULL)
		}

	case AgeParserT__3:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(41)
			p.Properties()
		}

	case AgeParserT__0:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(42)
			p.Arr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropertiesContext is an interface to support dynamic dispatch.
type IPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertiesContext differentiates from other interfaces.
	IsPropertiesContext()
}

type PropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesContext() *PropertiesContext {
	var p = new(PropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_properties
	return p
}

func (*PropertiesContext) IsPropertiesContext() {}

func NewPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesContext {
	var p = new(PropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_properties

	return p
}

func (s *PropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *PropertiesContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *PropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterProperties(s)
	}
}

func (s *PropertiesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitProperties(s)
	}
}

func (s *PropertiesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitProperties(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Properties() (localctx IPropertiesContext) {
	localctx = NewPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AgeParserRULE_properties)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(AgeParserT__3)
		}
		{
			p.SetState(46)
			p.Pair()
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AgeParserT__1 {
			{
				p.SetState(47)
				p.Match(AgeParserT__1)
			}
			{
				p.SetState(48)
				p.Pair()
			}

			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(54)
			p.Match(AgeParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(AgeParserT__3)
		}
		{
			p.SetState(57)
			p.Match(AgeParserT__4)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgeParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AgeParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(AgeParserSTRING)
	}
	{
		p.SetState(61)
		p.Match(AgeParserT__5)
	}
	{
		p.SetState(62)
		p.Value()
	}

	return localctx
}

// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgeParserRULE_arr
	return p
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgeParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.EnterArr(s)
	}
}

func (s *ArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgeListener); ok {
		listenerT.ExitArr(s)
	}
}

func (s *ArrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgeVisitor:
		return t.VisitArr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgeParser) Arr() (localctx IArrContext) {
	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AgeParserRULE_arr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(AgeParserT__0)
		}
		{
			p.SetState(65)
			p.Value()
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AgeParserT__1 {
			{
				p.SetState(66)
				p.Match(AgeParserT__1)
			}
			{
				p.SetState(67)
				p.Value()
			}

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(73)
			p.Match(AgeParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(AgeParserT__0)
		}
		{
			p.SetState(76)
			p.Match(AgeParserT__2)
		}

	}

	return localctx
}
