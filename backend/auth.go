package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	defaultUsername = "user"
	defaultPassword = "supersecurepassword"

	defaultTokenSecret = "supersecuresecret"
	defaultTokenIssuer = "oreo"

	tokenTypeBearer  = "Bearer"
	tokenTypeRefresh = "Refresh"
)

func areCredentialsValid(username, password string) bool {
	return username == defaultUsername && password == defaultPassword
}

type tokenClaims struct {
	jwt.StandardClaims
	TokenType string `json:"typ"`
}

type token struct {
	value     string
	expiresAt time.Time
}

func issueTokens(username string) (accessToken, refreshToken token, err error) {
	now := time.Now()

	accessToken.expiresAt = now.Add(1 * time.Minute)

	accessToken.value, err = jwt.
		NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			StandardClaims: jwt.StandardClaims{
				Subject:   username,
				Issuer:    defaultTokenIssuer,
				IssuedAt:  now.Unix(),
				ExpiresAt: accessToken.expiresAt.Unix(),
			},
			TokenType: tokenTypeBearer,
		}).
		SignedString([]byte(defaultTokenSecret))
	if err != nil {
		return token{}, token{}, err
	}

	refreshToken.expiresAt = now.Add(10 * time.Minute)

	refreshToken.value, err = jwt.
		NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			StandardClaims: jwt.StandardClaims{
				Subject:   username,
				Issuer:    defaultTokenIssuer,
				IssuedAt:  now.Unix(),
				ExpiresAt: refreshToken.expiresAt.Unix(),
			},
			TokenType: tokenTypeRefresh,
		}).
		SignedString([]byte(defaultTokenSecret))
	if err != nil {
		return token{}, token{}, err
	}

	return accessToken, refreshToken, nil
}

func parseAndVerifyToken(t string) (tokenClaims, error) {
	var claims tokenClaims

	if _, err := jwt.ParseWithClaims(t, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(defaultTokenSecret), nil
	}); err != nil {
		return tokenClaims{}, err
	}

	return claims, nil
}
