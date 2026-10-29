package main

type contextKey string

const (
	isAuthenticatedContextKey     = contextKey("isAuthenticated")
	authenticatedUserIDContextKey = contextKey("authenticatedUserID")
	flashContextKey               = contextKey("flash")
	redirectURIContextKey         = contextKey("redirectURI")
)

// Implement fmt.Stringer
func (key contextKey) String() string {
	return string(key)
}
