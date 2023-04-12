package config

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type Jwt struct {
	Key        string `yaml:"key"`
	ExpireTime int64  `yaml:"expire-time"` // 过期时间，单位分钟
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Jwt) Valid() {
	biz.IsTrue(j.Key != "", "config.yml之 [jwt.key] 不能为空")
	biz.IsTrue(j.ExpireTime != 0, "config.yml之 [jwt.expire-time] 不能为空")
}

func (j *Jwt) Generate(claims Claims) string {

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims.RegisteredClaims.ExpiresAt = &jwt.NumericDate{Time: expireTime}
	claims.RegisteredClaims.NotBefore = &jwt.NumericDate{Time: nowTime}
	claims.RegisteredClaims.IssuedAt = &jwt.NumericDate{Time: nowTime}
	claims.RegisteredClaims.Issuer = "chat"

	obj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := obj.SignedString(j.Key)
	if err != nil {
		return ""
	}
	return token
}

func (j *Jwt) GenerateToken(userId int) string {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: strconv.Itoa(userId),
		},
	}
	return j.Generate(claims)
}

func (j *Jwt) Parse(str string) (*jwt.Token, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(str, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.Key, nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return token, nil
	}
	return nil, errors.New("invalid token")
}
