package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zovdev1/mini-app-project/internal/entites"
	"github.com/zovdev1/mini-app-project/internal/repo"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
	auth "github.com/zovdev1/mini-app-project/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserCase struct {
	repo repo.UseRepo
	auth auth.TokenProvider
}

func New(repo repo.UseRepo, auth auth.TokenProvider) *UserCase {
	return &UserCase{
		repo: repo,
		auth: auth,
	}
}

func (u *UserCase) LogInUser(user dto.SignInInput) (string, error) {
	input, err := u.repo.FindByUser(user)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(input.Password), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("invalid email or password : %w", err)
	}

	token, err := u.auth.GenerateToken(input)

	if err != nil {
		return "", fmt.Errorf("UserUseCase - auth - u.auth.GenerateToken: %w", err)
	}

	return token, nil
}

func (u *UserCase) User(input dto.RegisterInput) (uuid.UUID, error) {

	_, err := u.repo.FindByUser(dto.SignInInput{ // полное хуйня
		Email: input.Email,
	})

	if err == nil {
		return uuid.UUID{}, fmt.Errorf("user with this email is already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 15)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to hash password: %w", err)
	}

	now := time.Now()

	user := entites.User{
		ID:        uuid.Must(uuid.NewV7()),
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = user.Validate()

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("validation failed: %w", err)
	}

	newUser, err := u.repo.Create(user)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to save user: %w", err)
	}

	return newUser, nil
}
