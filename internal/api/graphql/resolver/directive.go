package resolvers

import (
	"context"
	"strings"

	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// @role(roles: [String!]!)
func RoleDirective(
	ctx context.Context,
	obj interface{},
	next gqlgen.Resolver,
	roles []string,
) (interface{}, error) {

	// ------------------------------------------------
	// 1) Must be authenticated
	// ------------------------------------------------
	user, ok := UserFromCtx(ctx)
	if !ok || user == nil {
		return nil, &gqlerror.Error{
			Message: "unauthenticated",
			Extensions: map[string]interface{}{
				"code": "UNAUTHENTICATED",
			},
		}
	}

	userRole := strings.ToLower(user.Role)

	// ------------------------------------------------
	// 2) No roles specified → any authenticated user
	// ------------------------------------------------
	if len(roles) == 0 {
		return next(ctx)
	}

	// ------------------------------------------------
	// 3) Role check
	// ------------------------------------------------
	for _, r := range roles {
		if userRole == strings.ToLower(r) {
			return next(ctx)
		}
	}

	// ------------------------------------------------
	// 4) Forbidden
	// ------------------------------------------------
	return nil, &gqlerror.Error{
		Message: "forbidden",
		Extensions: map[string]interface{}{
			"code":          "FORBIDDEN",
			"user_role":     user.Role,
			"requiredRoles": roles,
		},
	}
}
