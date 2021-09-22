package model

type Result struct {
	Type   string
	Value  string
	Length int
}

func NewResult(Type string, Value string, Length int) Result {
	return Result{Type, Value, Length}
}
