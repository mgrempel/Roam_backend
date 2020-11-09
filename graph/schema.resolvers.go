package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Roam/Roam_backend/graph/generated"
	"Roam/Roam_backend/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context, userID int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByUUID(ctx context.Context, uuid string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
