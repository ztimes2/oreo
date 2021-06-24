package main

type tokenIssuer struct {
	refreshTokens map[string]struct{}
	secret string
}

func newTokenIssuer(secret string) *tokenIssuer {
	return &tokenIssuer{
		refreshTokens: make(map[string]struct{}),
		secret: secret,
	}
}

func (ti *tokenIssuer) parse(token string) (*jwt.Token , error) {
	return nil, errors.New("not implemented")
}

func (ti *tokenIssuer) issue(username string) (*oauth2.Token, error) {
	// TODO generate access token as jwt with:
	//  - the given username as subject
	//  - 1 minute expiry duration
	//  - the given secret as signature
	
	// TODO generate refresh token as jwt with:
	//  - the given username as subject
	//  - 10 minute expiry duration
	//  - the given secret as signature

	// TODO store refresh token in the map

	// TODO return tokens
	return nil, errors.New("not implemented")
}

func (ti *tokenIssuer) refresh(token string) (*oauth2.Token, error) {
	// TODO parse the given token
	
	// TODO validate the token
	
	// TODO lookup the token in the map 

	// TODO if token is invalid or not found in the map, then return error 

	// TODO if token is valid and found in the map, issue and return new tokens

	// TODO invalidate the incoming token by removing it from the map
	return nil, errors.New("not implemented")
}
