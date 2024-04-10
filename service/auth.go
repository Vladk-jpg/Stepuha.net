package service

import (
	"Stepuha.net/entities"
	"Stepuha.net/infrastructure"
	"crypto/sha512"
	"fmt"
)

const saltCrypto = "fdd12322dkanav22"

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

func generateHashedPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(saltCrypto)))
}
