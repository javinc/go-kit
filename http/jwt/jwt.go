package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// SigKey contains JWT signature key.
var SigKey = "goto-at-chiligarlic"

// Claims custom auth map claim.
type Claims struct {
	UID uint  `json:"uid"` // Get ID
	Lvl uint8 `json:"lvl"` // Auth Level
	jwt.StandardClaims
}

// New creates access token.
func New(c Claims) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(SigKey))
}

// Parse return id from token.
func Parse(token string) (c Claims, err error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SigKey), nil
	})

	if t == nil || !t.Valid {
		err = errors.New("token is not valid")
		return
	}

	tc, ok := t.Claims.(*Claims)
	if !ok || tc == nil {
		err = errors.New("token have empty claims")
		return
	}

	return *tc, nil
}

// ParseFrom http header Authorization request.
func ParseFrom(h http.Header) (c Claims, err error) {
	// Get token from Authorization header.
	t, err := pluckTokenFrom(h.Get("Authorization"))
	if err != nil {
		return
	}

	return Parse(t)
}

// plucks toke from a valid authorization bearer header.
func pluckTokenFrom(header string) (token string, err error) {
	// Get authorization header.
	header = strings.TrimSpace(header)
	if header == "" {
		err = errors.New("empty authorization header")
		return
	}

	// Should contain only Bearer and Token.
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		err = errors.New("invalid authorization header")
		return
	}

	// Check for bearer existence.
	if strings.ToUpper(parts[0]) != "BEARER" {
		err = errors.New("no bearer on authorization header")
		return
	}

	if strings.TrimSpace(parts[1]) == "" {
		err = errors.New("empty bearer token")
		return
	}

	return parts[1], nil
}
