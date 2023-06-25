package repository

import (
	"GitHubTrending/model"
	"GitHubTrending/model/req"
	"context"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
}
