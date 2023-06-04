package service

import (
	"encoding/base64"
	"fmt"
	"quest/global"
	"quest/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const IssuerName = "Skye>Quest.Auth"

func AuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ctx.Request.Header.Get("Authorization")
		if code == "" {
			global.ReturnError(ctx, global.Errors.NotLoginError)
			return
		} else if strings.Contains(code, " ") {
			code = code[strings.Index(code, " ")+1:]
		}

		info := jwt.MapClaims{}
		// 密钥
		secret := global.GetString("token.secret")
		token, err := jwt.ParseWithClaims(code, &info, func(token *jwt.Token) (interface{}, error) {
			key, err := base64.StdEncoding.DecodeString(secret)
			return key, err
		})
		if err != nil {
			global.ReturnError(ctx, global.Errors.TokenNotAvailableError)
			return
		}
		if !token.Valid {
			global.ReturnError(ctx, global.Errors.TokenInvalidError)
			return
		}
		iss, err := info.GetIssuer()
		if err != nil {
			global.ReturnError(ctx, global.Errors.TokenNotAvailableError)
			return
		}
		if iss != IssuerName {
			global.ReturnError(ctx, global.Errors.TokenIllegalError)
			return
		}
		sub, err := info.GetSubject()
		if err != nil {
			global.ReturnError(ctx, global.Errors.TokenNotAvailableError)
			return
		}
		subs := strings.Split(sub, "@")
		uid, err := strconv.ParseInt(subs[1], 10, 64)
		if err != nil {
			global.ReturnError(ctx, global.Errors.TokenNotAvailableError)
			return
		}
		user := model.User{
			Id:   uid,
			Name: subs[0],
		}
		ctx.Set("user", user)
	}
}

func GenerateToken(user *model.User) (string, int64, error) {
	// 密钥
	secret := global.GetString("token.secret")
	// 有效小时
	expTime := global.GetInt("token.exp")
	// 过期时间
	exp := time.Now().Add(time.Hour * time.Duration(expTime)).Unix()
	tc := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": exp,
			"iss": IssuerName,
			"sub": fmt.Sprintf("%s@%v", user.Name, user.Id),
		},
	)
	key, err := base64.StdEncoding.DecodeString(secret)
	token, err := tc.SignedString(key)
	return token, exp, err
}

func ValidateToken(code string) (bool, int, string, error) {
	info := jwt.MapClaims{}
	// 密钥
	secret := global.GetString("token.secret")
	token, err := jwt.ParseWithClaims(code, &info, func(token *jwt.Token) (interface{}, error) {
		key, err := base64.StdEncoding.DecodeString(secret)
		return key, err
	})
	if err != nil {
		return false, 0, "", err
	}
	if !token.Valid {
		return false, 0, "", nil
	}
	sub, err := info.GetSubject()
	if err != nil {
		return false, 0, "", err
	}
	subs := strings.Split(sub, "@")
	uid, err := strconv.Atoi(subs[1])
	if err != nil {
		return false, 0, "", err
	}
	return true, uid, subs[0], nil
}
