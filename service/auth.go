package service

import (
	"fmt"
	"quest/common"
	"quest/config"
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
			common.ReturnError(ctx, common.Errors.NotLoginError)
			return
		}

		info := jwt.MapClaims{}
		// 密钥
		secret := config.GetString("token.secret")
		token, err := jwt.ParseWithClaims(code, &info, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			common.ReturnError(ctx, common.Errors.TokenNotAvailableError)
			return
		}
		if !token.Valid {
			common.ReturnError(ctx, common.Errors.TokenInvalidError)
			return
		}
		iss, err := info.GetIssuer()
		if err != nil {
			common.ReturnError(ctx, common.Errors.TokenNotAvailableError)
			return
		}
		if iss != IssuerName {
			common.ReturnError(ctx, common.Errors.TokenIllegalError)
			return
		}
		sub, err := info.GetSubject()
		if err != nil {
			common.ReturnError(ctx, common.Errors.TokenNotAvailableError)
			return
		}
		subs := strings.Split(sub, "@")
		uid, err := strconv.Atoi(subs[1])
		if err != nil {
			common.ReturnError(ctx, common.Errors.TokenNotAvailableError)
			return
		}
		ctx.Set("token", subs[0])
		ctx.Set("uid", uid)
	}
}

func GenerateToken(user *model.User) (string, int64, error) {
	// 密钥
	secret := config.GetString("token.secret")
	// 有效小时
	expTime := config.GetInt("token.exp")
	// 过期时间
	exp := time.Now().Add(time.Hour * time.Duration(expTime)).Unix()
	tc := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": exp,
			"iss": IssuerName,
			"sub": fmt.Sprintf("%s@%v", user.Name, user.Id),
		},
	)
	token, err := tc.SignedString([]byte(secret))
	return token, exp, err
}

func ValidateToken(code string) (bool, int, string, error) {
	info := jwt.MapClaims{}
	// 密钥
	secret := config.GetString("token.secret")
	token, err := jwt.ParseWithClaims(code, &info, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
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
