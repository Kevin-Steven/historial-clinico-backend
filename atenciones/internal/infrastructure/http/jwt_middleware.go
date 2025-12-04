package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// NewAuthMiddleware crea un middleware que valida un JWT en el header Authorization: Bearer <token>.
// Usa el mismo secreto que el servicio de Auth.
func NewAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if jwtSecret == "" {
			// Si no hay secreto configurado, dejamos pasar (modo desarrollo)
			c.Next()
			return
		}

		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token requerido"})
			return
		}

		tokenStr := strings.TrimPrefix(h, "Bearer ")

		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		c.Next()
	}
}
