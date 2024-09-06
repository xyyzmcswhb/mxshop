package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint   `json:"id"`
	NickName    string `json:"nick_name"`
	AuthorityId uint   `json:"authority_id"` //是否为admin角色
	jwt.StandardClaims
}
