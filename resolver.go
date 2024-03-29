package GoGraphQLAPITesting

import (
	"context"

	model "github.com/PatriQ94/GoGraphQLAPITesting/models"
)

//User cache, should change with a real database
var usersCache []*model.User

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
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		CreatedAt: 39,
		UpdatedAt: 50}

	//Assign next ID, in real scenario should be setted by the database
	if len(usersCache) == 0 {
		nUser.ID = 1
	} else {
		var minID int = usersCache[0].ID
		for _, item := range usersCache {
			if item.ID > minID {
				minID = item.ID
			}
		}
		nUser.ID = minID + 1
	}
	//Add new user to the cache
	usersCache = append(usersCache, nUser)
	return nUser, nil
}
func (r *mutationResolver) RemoveUser(ctx context.Context, id int) (bool, error) {
	//If cache is empty, then return false
	if len(usersCache) == 0 {
		return false, nil
	}

	var removed bool = false
	for index, item := range usersCache {
		if item.ID == id {
			usersCache = append(usersCache[:index], usersCache[index+1:]...)
			removed = true
		}
	}

	if removed {
		return true, nil
	}
	return false, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return usersCache, nil
}
