package services

import (
	"errors"
	"strings"
	"time"

	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/hay-kot/homebox/backend/pkgs/hasher"
)

type GroupService struct {
	repos *repo.AllRepos
}

func (svc *GroupService) Get(ctx Context) (repo.Group, error) {
	return svc.repos.Groups.GroupByID(ctx.Context, ctx.GID)
}

func (svc *GroupService) UpdateGroup(ctx Context, data repo.GroupUpdate) (repo.Group, error) {
	if data.Name == "" {
		data.Name = ctx.User.GroupName
	}

	if data.Currency == "" {
		return repo.Group{}, errors.New("currency cannot be empty")
	}

	data.Currency = strings.ToLower(data.Currency)

	return svc.repos.Groups.GroupUpdate(ctx.Context, ctx.GID, data)
}

func (svc *GroupService) NewInvitation(ctx Context, uses int, expiresAt time.Time) (string, error) {
	token := hasher.GenerateToken()

	_, err := svc.repos.Groups.InvitationCreate(ctx, ctx.GID, repo.GroupInvitationCreate{
		Token:     token.Hash,
		Uses:      uses,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", err
	}

	return token.Raw, nil
}
