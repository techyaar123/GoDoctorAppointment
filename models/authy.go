package models

import "github.com/golang-jwt/jwt/v4"

type Authy struct{
	Username string `json:"username`
	Role string `json: "role`
	jwt.RegisteredClaims
}