package users

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Custom claims structure
type RegisterClaims struct {
	RejType string `json:",string"` //注册方式:	phone手机号注册 wxopen 微信开放平台登录 wxin 微信内登录 wxminip 微信小程序 pwd 账号密码注册
	Note    string //手机号 微信unionid 用户名
	Code    int64  //账密注册时的密码
	jwt.RegisteredClaims
}

func GetRegisterToken(secretKey string, t time.Time, seconds int64, RejType, Note string, Code int64) (string, error) {
	IssuedAt := jwt.NewNumericDate(t)
	claims := RegisterClaims{
		RejType: RejType,
		Note:    Note,
		Code:    Code,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t.Add(time.Duration(seconds) * time.Second)),
			IssuedAt:  IssuedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
