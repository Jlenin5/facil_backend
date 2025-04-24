package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// Definir clave de contexto para almacenar el usuario
type contextKey string

const UserContextKey contextKey = "user"

// Middleware para validar JWT y agregar el usuario al contexto
func JWTAuthMiddleware(secretKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "No autorizado", http.StatusUnauthorized)
				return
			}

			const bearerPrefix = "Bearer "
			if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
				http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
				return
			}
			tokenString := authHeader[len(bearerPrefix):]

			// Decodificar el token y obtener los claims
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
				return
			}

			// Extraer todos los datos del usuario desde los claims
			userData, ok := claims["user"].(map[string]interface{}) // Asegúrate de que el claim se llame "user"
			if !ok {
				http.Error(w, "Token inválido (sin datos de usuario)", http.StatusUnauthorized)
				return
			}

			// Guardar los datos del usuario en el contexto
			ctx := context.WithValue(r.Context(), UserContextKey, userData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}