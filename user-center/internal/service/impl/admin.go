package impl

import (
	"fmt"
	"hello/core/entity"
	"hello/core/mongo"
	"hello/core/utils"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Admin struct{}

// var tokenMap = map[string]UserInfo{}

type UserInfo struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Name         string   `json:"name"`
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
	Roles        []string `json:"roles"`
}

func (this_ *Admin) Login(reqData *entity.DtoAdminLogin) (interface{}, error) {
	fmt.Println("我来登陆了....")
	db := mongo.GetDb()
	defer db.Session.Close()
	var userInfo entity.UserInfo
	userInfoColl := db.C(userInfo.CollectionName())
	query := bson.M{"username": reqData.UserName}
	// 到数据库查询用户信息
	if err := userInfoColl.Find(query).One(&userInfo); err != nil {
		// 如果查询不到，则返回“不存在”
		if err == mgo.ErrNotFound {
			return nil, fmt.Errorf("用户不存在")
		} else { //其他错误信息直接返回错误信息
			return nil, err
		}
	}
	if utils.GetMd5(reqData.Password) != userInfo.Password {
		return nil, fmt.Errorf("密码错误")
	}
	// 密码正确
	// 生成一个令牌
	token := entity.Token{
		Id:       utils.GetUuid(),
		UserInfo: userInfo,
	}
	// 清空userinfo里的密码，以避免将密码返回客户端
	token.UserInfo.Password = ""
	tokenColl := db.C(token.CollectionName())
	// 删除所有旧token
	if _, err := tokenColl.RemoveAll(bson.M{"userInfo.username": reqData.UserName}); err != nil {
		return nil, err
	}
	if err := tokenColl.Insert(token); err != nil {
		return nil, err
	}
	token.UserInfo.Token = token.Id
	return token.UserInfo, nil
}

func (this_ *Admin) Logout(c *gin.Context) (interface{}, error) {

	fmt.Println("我来登出了....")

	tokenStr := c.Request.Header.Get("X-Token")

	db := mongo.GetDb()
	defer db.Session.Close()

	var token entity.Token
	tokenColl := db.C(token.CollectionName())

	if err := tokenColl.RemoveId(tokenStr); err != nil {
		return nil, err
	}

	return nil, nil
}

func (this_ *Admin) Info(reqData *entity.DtoAdminInfo) (interface{}, error) {
	return this_.Check(reqData.Token)
}

//通过Token获取用户信息
func (this_ *Admin) Check(tokenStr string) (*entity.UserInfo, error) {

	fmt.Println("我来获取用户信息了....")

	db := mongo.GetDb()
	defer db.Session.Close()

	var token entity.Token
	tokenColl := db.C(token.CollectionName())

	if err := tokenColl.FindId(tokenStr).One(&token); err != nil {
		if err == mgo.ErrNotFound {
			return nil, fmt.Errorf("令牌已过期")
		}
		return nil, err
	}

	return &token.UserInfo, nil
}

func (this_ *Admin) List(reqData *entity.DtoAdminUserList) (interface{}, error) {
	db := mongo.GetDb()
	defer db.Session.Close()
	var userInfo entity.UserInfo
	usrInfoColl := db.C(userInfo.CollectionName())
	var userInfoList []entity.UserInfo
	query := bson.M{}
	// or 查询
	if reqData.Keyword != "" {
		or := []bson.M{
			{"username": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"name": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"introduction": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
		}
		query["$or"] = or
	}
	var err error
	var count int
	// 查询分页的用户数据
	if err := usrInfoColl.Find(query).Skip((reqData.Page - 1) * reqData.Size).Limit(reqData.Size).All(&userInfoList); err != nil {
		return nil, err
	}
	// 查询总共数据条数
	if count, err = usrInfoColl.Find(query).Count(); err != nil {
		return nil, err
	}
	// 构造分页数据返回客户端
	pager := entity.Pager{
		Page:    reqData.Page,
		Size:    reqData.Size,
		Total:   count,
		Records: userInfoList,
	}
	return pager, nil
}

// 新增用户
func (this_ *Admin) Create(reqData *entity.UserInfo) (interface{}, error) {

	db := mongo.GetDb()
	defer db.Session.Close()

	var userInfo entity.UserInfo
	usrInfoColl := db.C(userInfo.CollectionName())
	// 检查用户是否存在
	if count, err := usrInfoColl.Find(bson.M{"username": reqData.Username}).Count(); err != nil {
		return nil, err
	} else if count > 0 {
		return nil, fmt.Errorf("用户%s已存在，请更换用户名", reqData.Username)
	}
	// 将客户端提交用户信息保存
	reqData.Id = bson.NewObjectId()
	if err := usrInfoColl.Insert(reqData); err != nil {
		return nil, err
	}

	return nil, nil
}

// 修改数据
func (this_ *Admin) Update(reqData *entity.UserInfo) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var userInfo entity.UserInfo
	usrInfoColl := db.C(userInfo.CollectionName())

	// 如果修改列表密码为空，则为原来密码
	if reqData.Password == "" {
		usrInfoColl.Find(bson.M{"_id": reqData.Id}).One(&userInfo)
		reqData.Password = userInfo.Password
	} else {
		// 对密码进行md5加密
		userInfo.Password = utils.GetMd5(reqData.Password)
		reqData.Password = userInfo.Password
	}

	// 对用户信息进行修改
	if err = usrInfoColl.Update(bson.M{"_id": reqData.Id}, reqData); err != nil {
		return nil, err
	}

	return nil, err
}

// 删除数据
func (this_ *Admin) Delete(reqData *entity.UserInfo) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var userInfo entity.UserInfo
	usrInfoColl := db.C(userInfo.CollectionName())

	if err = usrInfoColl.RemoveId(reqData.Id); err != nil {
		return nil, err
	}

	return nil, err
}
