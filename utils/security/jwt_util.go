package security

import (
	"fmt"
	"mnc/config"
	"mnc/model"
	exceptions "mnc/utils/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var blacklistedTokens = make(map[string]struct{})

func CreateAccessToken(user model.Customer) (string, error) {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	fmt.Printf("%v %v", ss, err)
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %v", err)
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, _ := config.NewConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != (*jwt.SigningMethodHMAC)(cfg.JwtSigningMethod) {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	if isTokenBlacklisted(tokenString) {
		return nil, fmt.Errorf("token has been invalidated")
	}
	return claims, nil
}

func InvalidateToken(tokenString string) {
	blacklistedTokens[tokenString] = struct{}{}
}

func isTokenBlacklisted(tokenString string) bool {
	_, blacklisted := blacklistedTokens[tokenString]
	return blacklisted
}
