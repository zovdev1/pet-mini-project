package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/zovdev1/mini-app-project/internal/entites"
)

type TokenJwt struct {
	secretKey     []byte
	tokenDuration time.Duration
}

func NewTokenJwt(secretKey string, tokenDuration time.Duration) TokenProvider {
	return &TokenJwt{
		secretKey:     []byte(secretKey),
		tokenDuration: tokenDuration,
	}
}

func (j *TokenJwt) GenerateToken(user *entites.User) (string, error) {

	claims := jwt.MapClaims{
		"exp": time.Now().Add(j.tokenDuration).Unix(),
		"iat": time.Now().Unix(),

		"sub":   user.ID.String(),
		"email": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secretKey)
}

func (j *TokenJwt) ValidateToken(tokenString string) (*entites.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		subVal, ok := claims["sub"]
		if !ok {
			return nil, errors.New("sub claim missing")
		}

		subStr, ok := subVal.(string)
		if !ok {
			return nil, errors.New("sub claim is not a string")
		}

		userID, err := uuid.Parse(subStr)
		if err != nil {
			return nil, fmt.Errorf("invalid uuid in sub: %w", err)
		}

		email, _ := claims["email"].(string)

		user := entites.User{
			ID:    userID,
			Email: email,
		}

		return &user, nil
	}

	return nil, errors.New("invalid token")
}
