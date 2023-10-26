package usecase

import (
	"fmt"
	"mnc/utils/security"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
	Logout(tokenString string) error
}

type authUseCase struct {
	usecase CustomerUseCase
}

func (a *authUseCase) Login(username, password string) (string, error) {
	user, err := a.usecase.FindByCustomerUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username and password or user is not active")
	}

	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}

func (a *authUseCase) Logout(tokenString string) error {
	security.InvalidateToken(tokenString)
	return nil
}

func NewAuthUseCase(usecase CustomerUseCase) AuthUseCase {
	return &authUseCase{usecase: usecase}
}
