package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	return service.UserCreate(ctx, input)
}

func (r *mutationResolver) UserUpdate(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	return service.UserUpdate(ctx, input)
}

func (r *mutationResolver) UserDelete(ctx context.Context, id int) (string, error) {
	return service.UserDelete(ctx, id)
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	return service.UserGetByID(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return service.UserGetAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
