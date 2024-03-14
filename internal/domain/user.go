package domain

import (
	"context"
	"errors"

	"github.com/Peterwmoss/LiCa/internal/auth"
	"github.com/Peterwmoss/LiCa/internal/database"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type (
	User struct {
		id    int
		Email string
	}

	UserService interface {
		GetOrCreate(token string) (*User, error)
		Get(token string) (*User, error)
		Create(email string) error
	}

	userService struct {
		db  *bun.DB
		ctx context.Context
	}
)

func NewUserService(db *bun.DB, ctx context.Context) UserService {
	return &userService{db, ctx}
}

func (svc userService) Get(token string) (*User, error) {
	userInfo, err := auth.GetUserInfo(token)
	if err != nil {
		return nil, err
	}

  log.Info().Msgf("Userinfo: %v", userInfo)

	user, err := svc.getIfExists(userInfo.Email)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to check if email exists: %s", userInfo.Email)
		return nil, err
	}

	if user == nil {
		err = errors.New("user not found: " + userInfo.Email)
		log.Error().Err(err).Send()
		return nil, err
	}

	return user, nil
}

func (svc userService) GetOrCreate(token string) (*User, error) {
	userInfo, err := auth.GetUserInfo(token)
	if err != nil {
		return nil, err
	}

	user, err := svc.getIfExists(userInfo.Email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get from database")
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	if err := svc.Create(userInfo.Email); err != nil {
    log.Error().Err(err).Msgf("Failed to create user with email: %s", userInfo.Email)
		return nil, err
	}

	user, err = svc.get(userInfo.Email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get created user from database")
		return nil, err
	}
	return user, nil
}

func (svc userService) Create(email string) error {
	_, err := svc.db.NewInsert().
		Model(&database.User{Email: email}).
		Exec(svc.ctx)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to create user with email: %s", email)
		return err
	}

	return nil
}

func (svc userService) getIfExists(email string) (*User, error) {
	exists, err := svc.emailExists(email)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, nil
	}

	return svc.get(email)
}

func (svc userService) get(email string) (*User, error) {
	dbUser := database.User{
		Email: email,
	}

	err := svc.db.NewSelect().
		Model(&dbUser).
		Scan(svc.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user from db")
		return nil, err
	}

	return &User{
		id:    dbUser.Id,
		Email: dbUser.Email,
	}, nil
}

func (svc userService) emailExists(email string) (bool, error) {
	user := database.User{
		Email: email,
	}

	exists, err := svc.db.NewSelect().
		Model(&user).
		Exists(svc.ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check if user exists:")
		return false, err
	}

	return exists, nil
}
