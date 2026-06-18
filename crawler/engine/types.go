package engine

import (
	"fmt"

	"github.com/javahongxi/golab/crawler/config"
)

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args any)
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload any
}

func GetPayload[T any](item *Item) (T, error) {
	var result T
	if item.Payload == nil {
		return result, nil
	}
	p, ok := item.Payload.(T)
	if !ok {
		return result, fmt.Errorf("payload type mismatch: expected %T, got %T", result, item.Payload)
	}
	return p, nil
}

func SetPayload[T any](item *Item, value T) {
	item.Payload = value
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args any) {
	return config.NilParser, nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args any) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
