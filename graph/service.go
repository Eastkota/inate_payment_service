package schema

import (
	"github.com/graphql-go/graphql"
)

var Service = graphql.NewObject(graphql.ObjectConfig{
	Name: "Service",
	Fields: graphql.Fields{
		"name":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"version": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"schema":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
	},
})
