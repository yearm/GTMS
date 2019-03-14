package controller

import (
	"github.com/astaxie/beego/context"
)

type Request struct {
	User    *SessionJson
	Context *context.Context
}

type SessionJson struct {
	AccessToken string `json:"accessToken"`
	UserInfo    `json:"userInfo"`
}

type UserInfo struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
