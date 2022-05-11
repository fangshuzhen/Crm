package impl

// "github.com/globalsign/mgo/bson"
import (
	"hello/core/entity"
	"hello/core/mongo"

	"github.com/globalsign/mgo/bson"
)

type Hello struct{}

// func (this_ *Hello) UserList(reqData *entity.DtoHello) (interface{}, error) {

// 	users := []entity.User{
// 		{Id: bson.NewObjectId(), Name: "小明1", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明2", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明3", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明4", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明4", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明4", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明4", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明4", UserName: "xiaoming", Password: "123456"},
// 		{Id: bson.NewObjectId(), Name: "小明4", UserName: "xiaoming", Password: "123456"},
// 	}

// 	return users, nil
// }

func (this_ *Hello) UserList(reqData *entity.DtoHello) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var user entity.User

	userColl := db.C(user.CollectionName())

	users := make([]entity.User, 10)

	if err = userColl.Find(nil).All(&users); err != nil {
		return nil, err
	}

	// fmt.Println(users)
	return users, nil
}

// 添加数据
func (this_ *Hello) InsertUser(reqData *entity.DtoUser) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var user entity.User
	userColl := db.C(user.CollectionName())

	user.Id = bson.NewObjectId()
	user.Username = reqData.UserName
	user.Password = reqData.Password
	user.Name = reqData.Name

	if err = userColl.Insert(user); err != nil {
		return nil, err
	}

	return nil, err
}

// 删除数据
func (this_ *Hello) DeleteUser(reqData *entity.DtoFindUser) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var user entity.User
	userColl := db.C(user.CollectionName())

	user.Name = reqData.Name

	if err = userColl.Remove(bson.M{"name": user.Name}); err != nil {
		return nil, err
	}

	return nil, err
}

// 查找数据
func (this_ *Hello) FindUser(reqData *entity.DtoFindUser) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var user entity.User

	userColl := db.C(user.CollectionName())

	users := make([]entity.User, 10)

	if err = userColl.Find(bson.M{"name":reqData.Name}).All(&users); err != nil {
		return nil, err
	}

	// fmt.Println(users)
	return users, nil
}


// 修改数据
func (this_ *Hello) EditUser(reqData *entity.DtoUser) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var user entity.User
	userColl := db.C(user.CollectionName())

	user.Username = reqData.UserName
	user.Password = reqData.Password

	if err = userColl.Update(bson.M{"name":reqData.Name},bson.M{"$set": bson.M{"password": reqData.Password,"username":reqData.UserName}}); err != nil {
		return nil, err
	}

	return nil, err
}