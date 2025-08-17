package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Handler.ServeHTTP(writer, request)
}

func (middleware AuthMiddleware) Middleware(role string, next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var key string = os.Getenv("AUTH_SECRET_KEY")
		var jwtKey = []byte(key)
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header format is invalid", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		if !isRoleAllowed(claims["role"].(string), role) {
			http.Error(w, "Forbidden: You don't have permission to access this resource", http.StatusForbidden)
			return
		}
		w.Header().Add("role", claims["role"].(string))
		w.Header().Add("user_id", claims["user_id"].(string))
		next(w, r, ps)
	}

}

// Helper function to check if user role is allowed
func isRoleAllowed(userRole string, allowedRoles string) bool {

	return strings.EqualFold(userRole, allowedRoles)
}
