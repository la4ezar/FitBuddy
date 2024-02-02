package domain

import "github.com/la4ezar/FitBuddy/pkg/graphql"

var _ graphql.ResolverRoot = &RootResolver{}

type RootResolver struct{}

func NewRootResolver() *RootResolver {
	return &RootResolver{}
}

type mutationResolver struct {
	*RootResolver
}

// Mutation missing godoc
func (r *RootResolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct {
	*RootResolver
}

// Query missing godoc
func (r *RootResolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}
