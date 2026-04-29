// generate access and refresh tokens using jwt and verify them

package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateTokens(userID string) (accessToken string, refreshToken string, err error) {
	// Create the JWT claims, which includes the user ID and expiry time
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	// Create the access token
	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = accessTokenClaims.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	// Create the refresh token
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = refreshTokenClaims.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func RefreshTokens(refreshTokenString string) (newAccessToken string, newRefreshToken string, err error) {
	claims, err := VerifyToken(refreshTokenString)
	if err != nil {
		return "", "", err
	}

	return GenerateTokens(claims.UserID)
}
