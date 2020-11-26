package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Roam/Roam_backend/graph/generated"
	"Roam/Roam_backend/graph/model"
	"Roam/Roam_backend/graph/utilities"
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
	var user model.User
	err := r.DB.Where("uuid = ?", input.UUID).First(&user).Error
	if err != nil {
		return nil, err
	}

	//Create a new post
	post := model.Post{
		Title:   input.Title,
		Content: input.Content,
		User:    &user}
	//Update the database
	r.DB.Create(&post)

	return &post, nil
}

func (r *mutationResolver) AddFriendByID(ctx context.Context, uuid string, id int) (*model.User, error) {
	var user model.User
	var friend model.User

	//Find us
	err := r.DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Error finding user")
	}

	//Find the friend
	err = r.DB.First(&friend, id).Error
	if err != nil {
		return nil, fmt.Errorf("Error locating friend")
	}
	// TODO: Implement a system for pending friend requests
	//Add friends
	r.DB.Model(&user).Association("Friends").Append(&friend)

	utilities.ScrubUser(&friend)
	return &friend, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByUUID(ctx context.Context, uuid string) (*model.User, error) {
	var user model.User
	r.DB.Where("uuid = ?", uuid).First(&user)
	utilities.ScrubUser(&user)
	return &user, nil
}

func (r *queryResolver) LogIn(ctx context.Context, username string, password string) (*model.User, error) {
	var user model.User
	// TODO: Change this to use email as a sign in. Rescaffold for consistency
	r.DB.Where("user_name = ?", username).Where("password = ?", password).Find(&user)
	user.Password = ""
	return &user, nil
}

func (r *queryResolver) GetUserPostsByUUID(ctx context.Context, uuid string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserFriendPostsByUUID(ctx context.Context, uuid string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserFriendsByUUID(ctx context.Context, uuid string) ([]*model.User, error) {
	var user model.User

	//Get user's Id
	err := r.DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Error fetching user")
	}

	//Find the users friends
	r.DB.Model(&user).Association("Friends").Find(&user.Friends)

	for _, friend := range user.Friends {
		//This isn't good, but for now I'll deal with the overhead.
		r.DB.Model(&friend).Association("Posts").Find(&friend.Posts)
		utilities.ScrubUser(friend)
	}

	return user.Friends, nil
}

func (r *queryResolver) GetUserTreeByUUID(ctx context.Context, uuid string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Error fetching user")
	}
	//Populate our posts
	r.DB.Model(&user).Association("Posts").Find(&user.Posts)

	utilities.ScrubUser(&user)
	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Posts(ctx context.Context, userID int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
