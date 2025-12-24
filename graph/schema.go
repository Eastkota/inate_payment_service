package schema

import (
    "github.com/graphql-go/graphql"
)

var queryType *graphql.Object
var mutationType *graphql.Object

var schema graphql.Schema
var err error

func InitSchema(q *graphql.Object, m *graphql.Object) {
    queryType = q
    mutationType = m
    schema, err = graphql.NewSchema(graphql.SchemaConfig{
        Query:    queryType,
        Mutation: mutationType,
    })
}

func GetSchema() (*graphql.Schema, error) {
    return &schema, err
}