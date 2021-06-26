package main

import (
	"net/http"
)

func handleVerify(w http.ResponseWriter, r *http.Request) {
	at := readAccessToken(r)
	if at == "" {
		writeError(w, http.StatusUnauthorized, errorResponse{
			Description: "Authentication required.",
		})
		return
	}

	claims, err := parseAndVerifyToken(at)
	if err != nil || claims.TokenType != tokenTypeBearer {
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
	t := readRefreshToken(r)
	if t == "" {
		writeError(w, http.StatusUnauthorized, errorResponse{
			Description: "Invalid refresh token.",
		})
		return
	}

	claims, err := parseAndVerifyToken(t)
	if err != nil || claims.TokenType != tokenTypeRefresh {
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
