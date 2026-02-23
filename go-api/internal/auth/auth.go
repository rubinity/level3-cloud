package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//sructure that keps token, namespace (as user id) and expiration time
type Tokens struct {
	token string
	JTIAcc   string
	ExpAcc   time.Time
	namespace   string
}

//issues an access token and sets a token structure
func IssueTokens(namespace string) (*Tokens, error) {
	now := time.Now().UTC()
	t := &Tokens{
		JTIAcc:   uuid.NewString(),
		namespace:  namespace,
		ExpAcc:   now.Add(1 * time.Hour),
	}

	acc := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        t.JTIAcc,
		Subject:   namespace,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(t.ExpAcc),
	})

	var err error
	t.token, err = acc.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return t, nil
}

//saves token id and namespace to redis db
func Persist(ctx context.Context, r *Redis, t *Tokens) error {
	if err := r.SetJTI(ctx, t.JTIAcc, t.namespace, t.ExpAcc); err != nil {
		return err
	}
	return nil
}

//puts token into a cookie
func SetAuthCookies(c *gin.Context, t *Tokens) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", t.token, int(time.Until(t.ExpAcc).Seconds()), "/", "", false, true)
}

//clears cookie
func ClearAuthCookies(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("access_token", "", -1, "/", "", true, true)
}

func ParseAccess(tokenStr string) (*jwt.RegisteredClaims, error) {
	secret := os.Getenv("ACCESS_SECRET")
	return parseWithSecret(tokenStr, secret)
}

// parser 
// checks if a secret exists
// creates a new parser with allowed methods
// checks signing methods
// checks if valid
// checks if expired
func parseWithSecret(tokenStr, secret string) (*jwt.RegisteredClaims, error) {
	if secret == "" {
		return nil, errors.New("jwt secret not configured")
	}

	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
// takes a claims structure to fill, it decodes the token, verifies signature, populates claim structure
	token, err := parser.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Extra safety: ensure HMAC family
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
