package driver

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type JWT struct {
	Uuid                  uuid.UUID `json:"-"`
	AccessTokenString     string    `json:"access_token_string"`
	RefreshTokenString    string    `json:"refresh_token_string"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
}

// JWTClaims defines the structure of the JWT claims.
// TODO: add client_id latter
type JWTClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

// GenerateToken generates an access token and a refresh token.
func (j *JWT) GenerateToken(email string) (interface{}, error) {
	tokenUuid, _ := uuid.NewUUID()
	accessTokenLifetime, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFETIME_SEC"))
	accessTokenExpiresAt := time.Now().Add(time.Duration(accessTokenLifetime) * time.Second)
	accessTokenString, err := generateToken(tokenUuid, accessTokenExpiresAt, email)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshTokenLifetime, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION_SEC"))
	refreshTokenExpiresAt := time.Now().Add(time.Duration(refreshTokenLifetime) * time.Second)
	refreshTokenString, err := generateToken(tokenUuid, refreshTokenExpiresAt, email)
	if err != nil {
		return nil, err
	}

	return &JWT{
		Uuid:                  tokenUuid,
		AccessTokenString:     accessTokenString,
		RefreshTokenString:    refreshTokenString,
		RefreshTokenExpiresAt: refreshTokenExpiresAt,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
	}, nil
}

// ValidateToken validates a token string and returns the claims if the token is valid.
func (j *JWT) ValidateToken(tokenString string) (interface{}, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithAudience(os.Getenv("APP_HOST")), jwt.WithIssuer(os.Getenv("APP_NAME")))

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// generateToken creates a token with a specified expiration duration.
func generateToken(uuid uuid.UUID, expiresAt time.Time, email string) (string, error) {
	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.String(),
			Issuer:    os.Getenv("APP_NAME"),
			Audience:  []string{os.Getenv("APP_HOST")},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		Email: email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
