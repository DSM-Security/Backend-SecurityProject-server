package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// test jwtSignKey
var testSignKey = []byte("TestForFasthttpWithJWT")

// Credential type
type userCredential struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	School            string `json:"school"`
	Nickname          string `json:"nickname"`
	HashedAccessToken string `json:"hashedAccessToken"`
	jwt.StandardClaims
}

type JwtPayload struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

// GetToken function
func GetTokenString(c *fiber.Ctx) ([]byte, error) {
	// Get token from request token
	jwt := c.Request().Header.Peek("Authorization")

	// Token length validation
	if len(jwt) == 0 {
		c.SendStatus(401)
		return nil, errors.New("Token cannot found")
	}

	// Return token with type []byte
	return jwt, nil
}

// Generate accessToken
func AccessToken(data JwtPayload) string {
	// Generate Token object
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Id:       data.Id,
		Nickname: data.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(), // 10 Mins
		},
	})

	// jwt Key
	jwtSignKey := []byte(GetSecretKey())

	// Sign token
	access, err := accessToken.SignedString(jwtSignKey)
	HandleErr(err)

	return access
}

// Validate token
func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	// Generate Credential object
	user := &userCredential{}

	// jwt Key
	jwtSignKey := []byte(GetSecretKey())

	// Parse token and validate
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	HandleErr(err)

	return token, user, err
}
