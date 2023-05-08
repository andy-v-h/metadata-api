package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

// AnnotationNamespace returns AnnotationNamespaceResolver implementation.
func (r *Resolver) AnnotationNamespace() AnnotationNamespaceResolver {
	return &annotationNamespaceResolver{r}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// StatusNamespace returns StatusNamespaceResolver implementation.
func (r *Resolver) StatusNamespace() StatusNamespaceResolver { return &statusNamespaceResolver{r} }

type annotationNamespaceResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type statusNamespaceResolver struct{ *Resolver }