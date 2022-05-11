package entity

import "github.com/globalsign/mgo/bson"

type UserInfo struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Username     string        `json:"username" bson:"username"`
	Password     string        `json:"password" bson:"password"`
	Name         string        `json:"name" bson:"name"`
	Avatar       string        `json:"avatar" bson:"avatar"`
	Introduction string        `json:"introduction" bson:"introduction"`
	Roles        []string      `json:"roles" bson:"roles"`
	Token        string        `json:"token" bson:"-"`
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Name     string        `json:"name" bson:"name"`
}

func (this_ *User) CollectionName() string {
	return "user"
}

// 用于返回集合名称
func (this_ *UserInfo) CollectionName() string {
	return "user_info"
}

// 用户登录令牌
type Token struct {
	Id       string   `json:"id" bson:"_id"`
	UserInfo UserInfo `json:"userInfo" bson:"userInfo"`
}

func (this_ *Token) CollectionName() string {
	return "token"
}

//获取用户列表的请求模板
type DtoAdminUserList struct {
	DtoPageList
}

//登陆请求数据模板
type DtoAdminLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type DtoAdminInfo struct {
	Token string `json:"token"`
}
