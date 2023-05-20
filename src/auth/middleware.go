package auth

import (
	"net/http"
	"strings"
)

// middleware pour valider et extraire le token  des requêtes
func MiddlewareJWT(jwtManager JWTManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Récupérer le token JWT du header Authorization
			authHeader := r.Header.Get("Authorization")
			token := extractJWTFromHeader(authHeader)

			// Vérifier si le token est valide
			if token != "" {
				userID, err := jwtManager.VerifyToken(token)
				if err == nil {
					// Ajouter l'ID utilisateur extrait à la requête
					r.Header.Set("UserID", userID)
				}
			}

			// Appel le gestionnaire suivant dans la chaîne
			next.ServeHTTP(w, r)
		})
	}
}

// extrait le token du header Authorization
func extractJWTFromHeader(header string) string {
	parts := strings.Split(header, "")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
