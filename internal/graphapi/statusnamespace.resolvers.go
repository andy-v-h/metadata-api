package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/x/gidx"
)

// StatusNamespaceCreate is the resolver for the statusNamespaceCreate field.
func (r *mutationResolver) StatusNamespaceCreate(ctx context.Context, input generated.CreateStatusNamespaceInput) (*StatusNamespaceCreatePayload, error) {
	// TODO: authz check here
	ns, err := r.client.StatusNamespace.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &StatusNamespaceCreatePayload{StatusNamespace: ns}, nil
}

// StatusNamespaceDelete is the resolver for the statusNamespaceDelete field.
func (r *mutationResolver) StatusNamespaceDelete(ctx context.Context, id gidx.PrefixedID, force bool) (*StatusNamespaceDeletePayload, error) {
	// TODO: authz check here
	statusCount, err := r.client.Status.Query().Where(status.StatusNamespaceID(id)).Count(ctx)
	if err != nil {
		return nil, err
	}

	if statusCount != 0 {
		if force {
			statusCount, err = r.client.Status.Delete().Where(status.StatusNamespaceID(id)).Exec(ctx)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("status namespace is in use and can't be deleted")
		}
	}

	err = r.client.StatusNamespace.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &StatusNamespaceDeletePayload{DeletedID: id, StatusDeletedCount: statusCount}, nil
}

// StatusNamespaceUpdate is the resolver for the statusNamespaceUpdate field.
func (r *mutationResolver) StatusNamespaceUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateStatusNamespaceInput) (*StatusNamespaceUpdatePayload, error) {
	// TODO: authz check here
	ns, err := r.client.StatusNamespace.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	ns, err = ns.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &StatusNamespaceUpdatePayload{StatusNamespace: ns}, nil
}
