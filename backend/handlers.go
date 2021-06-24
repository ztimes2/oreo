package main

import (
	"net/http"
)

type handler struct {
	ti *tokenIssuer
}

func (h *handler) handleVerify(w http.ResponseWriter, r *http.Request) {
	// TODO read access token either from cookie or header
	
	// TODO verify access token

	// TODO respond with error if access token is not valid

	// TODO respond with success if access token is valid
}

func (h *handler) handleSignIn(w http.ResponseWriter, r *http.Request) {
	// TODO read username and password from body

	// TODO verify username and password

	// TODO respond with error if credentials are not valid

	// TODO generate access and refresh tokens, and respond with them if credentials
	// are valid
}

func (h *handler) handleRefresh(w http.ResponseWriter, r *http.Request) {
	// TODO read refresh token from body
	
	// TODO verify refresh token

	// TODO respond with error if refresh token is not valid

	// TODO generate access and refresh tokens, and respond with them if refresh
	// token is valid
}
