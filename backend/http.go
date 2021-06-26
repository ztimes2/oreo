package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

const (
	headerContentType   = "Content-Type"
	headerAuthorization = "Authorization"

	contentTypeJSON = "application/json"
)

func writeJSON(w http.ResponseWriter, statusCode int, resp interface{}) {
	if resp == nil {
		w.WriteHeader(statusCode)
		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		writeUnexpectedError(w, err)
		return
	}

	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)
	w.Write(body)
}

func writeUnexpectedError(w http.ResponseWriter, err error) {
	logrus.Errorf("unexpected error: %v", err)

	body, _ := json.Marshal(errorResponse{
		Description: "Something went wrong. Please, try again later.",
	})

	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(body)
}

type errorResponse struct {
	Description string `json:"err_description"`
}

func writeError(w http.ResponseWriter, statusCode int, e errorResponse) {
	writeJSON(w, statusCode, e)
}

func writeTokens(w http.ResponseWriter, accessToken, refreshToken token) {
	writeJSON(w, http.StatusOK, oauth2.Token{
		AccessToken:  accessToken.value,
		Expiry:       accessToken.expiresAt,
		TokenType:    tokenTypeBearer,
		RefreshToken: refreshToken.value,
	})
}

func readAccessToken(r *http.Request) string {
	h := r.Header.Get(headerAuthorization)
	segments := strings.Split(h, " ")
	if len(segments) == 2 && segments[0] == tokenTypeBearer {
		return segments[1]
	}
	return ""
}

func readRefreshToken(r *http.Request) string {
	if r.PostForm == nil {
		r.ParseForm()
	}
	return r.PostFormValue("refresh_token")
}
