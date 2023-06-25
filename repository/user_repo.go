package repository

import (
	"GitHubTrending/model"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
