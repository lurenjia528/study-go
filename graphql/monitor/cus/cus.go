package cus

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"math"
	"strconv"
)

// Int is the GraphQL Integer type definition.
var Int64 = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Int64",
	Description: "int64",
	Serialize:  coerceInt64,
	ParseValue: coerceInt64,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			fmt.Println(valueAST)
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})

func coerceInt64(value interface{}) interface{} {
	switch value := value.(type) {
	case int64:
		if value < int64(math.MinInt32) || value > int64(math.MaxInt32) {
			return value
		}
		return value
	case *int64:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case uint64:
		if value > uint64(math.MaxInt32) {
			return value
		}
		return value
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}