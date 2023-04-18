package filter

import (
	"fmt"
)

const (
	DataTypeStr  = "string"
	DataTypeInt  = "int"
	DataTypeDate = "data"
	DataTypeBool = "bool"

	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "lte"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "gte"
	OperatorBetween       = "between"
	OperatorLike          = "like"
)

type Opptions struct {
	limit  int
	fields []Field
}

func NewOptions(limit int) *Opptions {
	return &Opptions{limit: limit}
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

type Options interface {
	Limit() int
	AddField(name string, operator string, value string, dtype string) error
	Fields() []Field
}

func (o *Opptions) Limit() int {
	return o.limit
}

func (o *Opptions) AddField(name string, operator string, value string, dtype string) error {
	err := validateOperator(operator)
	if err != nil {
		return err
	}

	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: operator,
		Type:     dtype,
	})

	fmt.Println(o.fields)
	return nil
}

func (o *Opptions) Fields() []Field {
	return o.fields
}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThan:
	case OperatorGreaterThanEq:
	case OperatorLike:
	case OperatorBetween:
	default:
		return fmt.Errorf("bad operator")
	}
	return nil
}
