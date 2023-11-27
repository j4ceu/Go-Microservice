package utils

import (
	"Go-Microservice/app/user/config"
	"Go-Microservice/app/user/models"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(48 * time.Hour)
	authConfig := config.LoadAuthConfig()
	claims := &JWTClaim{
		UserID: user.ID.String(),
		Email:  user.Email,
		Role:   user.Role.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(authConfig.Secret))
	return
}

func ValidateToken(signedToken string, role string) (*JWTClaim, error) {
	authConfig := config.LoadAuthConfig()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(authConfig.Secret), nil
		},
	)

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	if role != "" {
		if claims.Role != role {
			err = errors.New("you dont have authorized")
			return claims, err
		}
		return claims, err
	}
	return claims, err
}

func GetTokenClaims(ctx *gin.Context) *JWTClaim {
	claims := ctx.MustGet("claims").(*JWTClaim)
	return claims

}
