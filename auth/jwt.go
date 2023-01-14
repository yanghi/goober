package auth

import (
	"errors"
	"goober/goober"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type DefaultClaims struct {
	JwtUserClaims
	jwt.RegisteredClaims
}

type JwtUserClaims struct {
	Username string `json:"name"`
	// 用户id
	Uid int64 `json:"id"`
}

type JwtConfig struct {
	// 过期时间
	Expire time.Duration
	// 生成token的密钥
	SignedKey interface{}
}

var conf JwtConfig = JwtConfig{
	Expire:    time.Hour * 72,
	SignedKey: []byte("bdebea55-ec6f-41b5-8f29-0433509f00fb"),
}

type JwtUserAuth struct {
	JwtUserClaims
}

//生成token
func GenToken(username string, uid int64) (string, error) {
	c := DefaultClaims{
		JwtUserClaims{
			Username: username,
			Uid:      uid,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(conf.Expire)),
			Issuer:    "goober",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(conf.SignedKey)
}

func ParseToken(tokenString string) (*DefaultClaims, *goober.Error) {

	token, err := jwt.ParseWithClaims(tokenString, &DefaultClaims{}, func(token *jwt.Token) (interface{}, error) {
		return conf.SignedKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			goober.NewWithCode(goober.ErrTokenExpired)

			return nil, goober.NewWithCode(goober.ErrTokenExpired)
		}
	}
	if claims, ok := token.Claims.(*DefaultClaims); ok && token.Valid {
		return claims, nil
	} else {

		return nil, goober.NewWithCode(goober.ErrTokenInvalid)
	}
}

func GetUser(tokenString string) (*JwtUserClaims, *goober.Error) {
	claims, e := ParseToken(tokenString)
	if e != nil {
		return nil, e
	}

	return &JwtUserClaims{Username: claims.Username, Uid: claims.Uid}, nil
}
