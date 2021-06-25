package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func newRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestLogger(&middleware.DefaultLogFormatter{
			Logger: logrus.StandardLogger(),
		}),
		middleware.Recoverer,
	)

	r.Post("/signin", handleSignIn)
	r.Post("/refresh", handleRefresh)
	r.Post("/verify", handleVerify)

	return r
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	at := readAccessToken(r)

	claims, err := parseAndVerifyToken(at)
	if err != nil || claims.TokenType != tokenTypeBearer {
		logrus.Errorf("invalid access token: %v", err)
		writeError(w, http.StatusUnauthorized, errorResponse{
			Description: "Authentication required.",
		})
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func handleSignIn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if !areCredentialsValid(username, password) {
		writeError(w, http.StatusBadRequest, errorResponse{
			Description: "Invalid username or password.",
		})
		return
	}

	at, rt, err := issueTokens(username)
	if err != nil {
		writeUnexpectedError(w, err)
		return
	}

	writeTokens(w, at, rt)
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	refreshToken := readRefreshToken(r)

	claims, err := parseAndVerifyToken(refreshToken)
	if err != nil || claims.TokenType != tokenTypeRefresh {
		logrus.Errorf("invalid refresh token: %v", err)
		writeError(w, http.StatusBadRequest, errorResponse{
			Description: "Invalid refresh token.",
		})
		return
	}

	at, rt, err := issueTokens(claims.Subject)
	if err != nil {
		writeUnexpectedError(w, err)
		return
	}

	writeTokens(w, at, rt)
}
