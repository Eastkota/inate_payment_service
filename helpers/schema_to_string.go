package helpers

import (
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
)

func ConvertSchemaToString(schema *graphql.Schema) string {
	var sb strings.Builder

	// List of built-in GraphQL scalars
	builtIns := map[string]bool{
		"String":  true,
		"Int":     true,
		"Float":   true,
		"Boolean": true,
		"ID":      true,
	}

	for typeName, graphqlType := range schema.TypeMap() {
		// Ignore internal GraphQL types
		if strings.HasPrefix(typeName, "__") {
			continue
		}

		switch t := graphqlType.(type) {
		case *graphql.Object:
			sb.WriteString(fmt.Sprintf("type %s {\n", typeName))
			for fieldName, field := range t.Fields() {
				fieldTypeName := field.Type.String()
				if fieldType, ok := field.Type.(*graphql.NonNull); ok {
					fieldTypeName = fmt.Sprintf("%s!", fieldType.OfType.String())
				}

				// Handle arguments
				var args []string
				for _, arg := range field.Args {
					argTypeName := arg.Type.String()
					if argType, ok := arg.Type.(*graphql.NonNull); ok {
						argTypeName = fmt.Sprintf("%s!", argType.OfType.String())
					}
					args = append(args, fmt.Sprintf("%s: %s", arg.Name(), argTypeName))
				}

				argString := ""
				if len(args) > 0 {
					argString = fmt.Sprintf("(%s)", strings.Join(args, ", "))
				}

				sb.WriteString(fmt.Sprintf("  %s%s: %s\n", fieldName, argString, fieldTypeName))
			}
			sb.WriteString("}\n\n")

		case *graphql.InputObject:
			sb.WriteString(fmt.Sprintf("input %s {\n", typeName))
			for fieldName, field := range t.Fields() {
				fieldTypeName := field.Type.String()
				if fieldType, ok := field.Type.(*graphql.NonNull); ok {
					fieldTypeName = fmt.Sprintf("%s!", fieldType.OfType.String())
				}
				sb.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, fieldTypeName))
			}
			sb.WriteString("}\n\n")

		case *graphql.Enum:
			sb.WriteString(fmt.Sprintf("enum %s {\n", typeName))
			for _, value := range t.Values() {
				sb.WriteString(fmt.Sprintf("  %s\n", value.Name))
			}
			sb.WriteString("}\n\n")

		case *graphql.Interface:
			sb.WriteString(fmt.Sprintf("interface %s {\n", typeName))
			for fieldName, field := range t.Fields() {
				fieldTypeName := field.Type.String()
				if fieldType, ok := field.Type.(*graphql.NonNull); ok {
					fieldTypeName = fmt.Sprintf("%s!", fieldType.OfType.String())
				}
				sb.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, fieldTypeName))
			}
			sb.WriteString("}\n\n")

		case *graphql.Union:
			sb.WriteString(fmt.Sprintf("union %s = ", typeName))
			for i, ut := range t.Types() {
				if i > 0 {
					sb.WriteString(" | ")
				}
				sb.WriteString(ut.String())
			}
			sb.WriteString("\n\n")

		case *graphql.Scalar:
			// Only write custom scalars
			if !builtIns[typeName] {
				sb.WriteString(fmt.Sprintf("scalar %s\n\n", typeName))
			}
		}
	}

	return sb.String()
}
