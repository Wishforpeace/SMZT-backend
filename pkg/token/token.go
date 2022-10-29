package token

import (
	"SMZT/pkg/errno"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"SMZT/log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	jwtKey string

	ErrTokenInvalid = errors.New("the token is invalid")
	ErrTokenExpired = errors.New("the token is expired")
)

// getJwtKey get jwtKey.
func getJwtKey() string {
	if jwtKey == "" {
		jwtKey = viper.GetString("jwt_secret")
	}
	return jwtKey
}

// TokenPayload is a required payload when generates token.
type TokenPayload struct {
	Id        uint          `json:"id"`
	StudentID string        `json:"student_id"`
	Expired   time.Duration `json:"expired"` // 有效时间（nanosecond）
}

// TokenResolve means returned payload when resolves token.
type TokenResolve struct {
	Id        uint   `json:"id"`
	StudentID string `json:"student_id"`
	ExpiresAt int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

// GenerateToken generates token.
func (payload *TokenPayload) GenerateToken() (string, error) {
	claims := &TokenClaims{
		Id:        payload.Id,
		StudentID: payload.StudentID,
		ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
	}
	fmt.Println("ExpiresAt: ", time.Now().Unix()+int64(payload.Expired.Seconds()))
	encodedString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := encodedString.SignedString([]byte(getJwtKey()))
	if err != nil {
		return "", errno.ErrFormToken
	}

	return token, nil
}

// ResolveToken resolves token.
func ResolveToken(tokenStr string) (*TokenResolve, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		log.Error("Token parsing failed because of an internal error", zap.String("cause", err.Error()))
		return nil, err
	}

	if !token.Valid {
		log.Error("Token parsing failed; the token is invalid.")
		return nil, ErrTokenInvalid
	}

	t := TokenResolve{
		Id:        claims.Id,
		StudentID: claims.StudentID,
		ExpiresAt: claims.ExpiresAt,
	}
	fmt.Println("token resolve", t)
	return &t, nil
}
