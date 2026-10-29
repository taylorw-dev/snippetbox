package main

type contextKey string

const (
	isAuthenticatedContextKey     = contextKey("isAuthenticated")
	authenticatedUserIDContextKey = contextKey("authenticatedUserID")
	flashContextKey               = contextKey("flash")
	redirectPathContextKey        = contextKey("redirectPath")
)

// Implement fmt.Stringer
func (key contextKey) String() string {
	return string(key)
}
