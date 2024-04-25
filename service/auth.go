package service

import (
	"Stepuha.net/entities"
	"Stepuha.net/infrastructure"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	saltCrypto = "fdd12322dkanav22"
	signingKey = "drt#45ytu67yikhgv"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repos infrastructure.Authorization
}

func NewAuthService(repository infrastructure.Authorization) *AuthService {
	return &AuthService{repos: repository}
}

func (serv *AuthService) AddUser(user entities.User) (int, error) {
	user.Password = generateHashedPassword(user.Password)
	return serv.repos.AddUser(user)
}

func (serv *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := serv.repos.GetUser(username, generateHashedPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (serv *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func generateHashedPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltCrypto)))
}
