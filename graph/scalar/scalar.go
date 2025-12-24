package scalar

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"time"
)

var UUID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "UUID",
	Description: "The UUID scalar type represents a UUID as defined in RFC 4122.",
	Serialize: func(value interface{}) interface{} {
		switch v := value.(type) {
		case uuid.UUID:
			return v.String()
		case string:
			return v
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		if str, ok := value.(string); ok {
			id, err := uuid.Parse(str)
			if err == nil {
				return id
			}
		}
		return nil
	},
	ParseLiteral: func(astValue ast.Value) interface{} {
		if v, ok := astValue.(*ast.StringValue); ok {
			id, err := uuid.Parse(v.Value)
			if err == nil {
				return id
			}
		}
		return nil
	},
})

var Time = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Time",
	Description: "The Time scalar type represents time values in RFC3339 format",

	// Serialize outgoing time.Time to string
	Serialize: func(value interface{}) interface{} {
		if t, ok := value.(time.Time); ok {
			return t.Format(time.RFC3339)
		}
		return nil
	},

	// Parse value from GraphQL input variables
	ParseValue: func(value interface{}) interface{} {
		if str, ok := value.(string); ok {
			t, err := time.Parse(time.RFC3339, str)
			if err == nil {
				return t
			}
		}
		return nil
	},

	// Parse literal from GraphQL query (inline string)
	ParseLiteral: func(astValue ast.Value) interface{} {
		if v, ok := astValue.(*ast.StringValue); ok {
			t, err := time.Parse(time.RFC3339, v.Value)
			if err == nil {
				return t
			}
		}
		return nil
	},
})