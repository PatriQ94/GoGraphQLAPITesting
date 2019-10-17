package GoGraphQLAPITesting

import (
	"context"

	model "github.com/PatriQ94/GoGraphQLAPITesting/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var ret []*model.User

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	nUser := &model.User{
		ID:        "1",
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		CreatedAt: 39,
		UpdatedAt: 50,
	}

	ret = append(ret, nUser)

	return nUser, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//List of users to return

	return ret, nil
}
