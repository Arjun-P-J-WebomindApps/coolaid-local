package helpers

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// GetRequestedSubFields returns the list of field names requested
// by the client under a specific parent field in the GraphQL query.
//
// Example:
//
// Query:
//
//	filterProducts {
//	  self {
//	    companyName
//	    modelName
//	  }
//	}
//
// Calling:
//
//	GetRequestedSubFields(ctx, "self")
//
// Returns:
//
//	[]string{"companyName", "modelName"}
//
// This function belongs in the GraphQL delivery layer because
// it depends on GraphQL AST and request context.
func GetRequestedSubFields(ctx context.Context, parentField string) []string {
	opCtx := graphql.GetOperationContext(ctx)

	// Locate the selection set for the specified parent field
	var selection ast.SelectionSet

	for _, field := range graphql.CollectFieldsCtx(ctx, nil) {
		if field.Name == parentField {
			selection = field.Field.SelectionSet
			break
		}
	}

	if selection == nil {
		return nil
	}

	// Collect requested child fields
	collected := graphql.CollectFields(opCtx, selection, nil)

	fields := make([]string, 0, len(collected))
	for _, f := range collected {
		fields = append(fields, f.Name)
	}

	return fields
}
