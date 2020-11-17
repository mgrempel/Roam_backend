package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Roam/Roam_backend/graph/generated"
	"Roam/Roam_backend/graph/model"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	//Generate a UUID for the user
	id := uuid.New()
	//Create our new user
	newUser := model.User{
		UserName:    input.UserName,
		Password:    input.Password,
		Email:       input.Email,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Description: input.Description,
		UUID:        id.String()}

	// TODO: Need to ensure an account doesn't exist already.

	r.DB.Create(&newUser)

	return &newUser, nil
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

func (r *queryResolver) LogIn(ctx context.Context, username string, password string) (*model.User, error) {
	var user model.User
	// TODO: Change this to use email as a sign in. Rescaffold for consistency
	r.DB.Where("user_name = ?", username).Where("password = ?", password).Find(&user)
	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
