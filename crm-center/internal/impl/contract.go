package impl

import (
	"fmt"
	"hello/core/entity"
	"hello/core/mongo"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type Contract struct{}

// 文件上传
func (this_ *Contract) UploadContractFile(c *gin.Context) (interface{}, error) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "文件上传错误")
	}
	// 获取用户id
	contractId := c.DefaultPostForm("contract", "")
	// 上传文件保存路径
	// path := fmt.Sprintf("%s/%s", "E:/", file.Filename)
	basePath := "C:/Program Files/Go/src/github.com/hello/cmd/admin/web/vue-element-admin-master/src/assets/path/"
	// 将图片存放到前端静态文件夹/ 用户id / 文件名
	DTpath, _ := CreateDateDir(basePath, contractId)
	path := fmt.Sprintf("%s/%s", DTpath, file.Filename)

	c.SaveUploadedFile(file, path)

	result := map[string]string{
		"file": file.Filename,
	}
	return result, nil
}

//basePath是固定目录路径
func CreateDateDir(basePath string, folderName string) (dirPath, dataString string) {
	// folderName := time.Now().Format("2006-01-02")
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath, folderName
}

//数据展示
func (this_ *Contract) List(reqData *entity.DtoContractList) (interface{}, error) {
	db := mongo.GetDb()
	defer db.Session.Close()
	var contract entity.Contract
	contractColl := db.C(contract.CollectionName())
	var contractList []entity.Contract
	query := bson.M{}
	// or 查询
	if reqData.Keyword != "" {
		or := []bson.M{
			{"title": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"state": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
			{"desc": bson.M{"$regex": fmt.Sprintf("^.*%s.*", reqData.Keyword)}},
		}
		query["$or"] = or
	}
	var err error
	var count int
	// 查询分页的用户数据
	if err := contractColl.Find(query).Skip((reqData.Page - 1) * reqData.Size).Limit(reqData.Size).All(&contractList); err != nil {
		return nil, err
	}
	// 查询总共数据条数
	if count, err = contractColl.Find(query).Count(); err != nil {
		return nil, err
	}
	// 构造分页数据返回客户端
	pager := entity.Pager{
		Page:    reqData.Page,
		Size:    reqData.Size,
		Total:   count,
		Records: contractList,
	}
	return pager, nil
}

// 添加数据
func (this_ *Contract) Create(reqData *entity.Contract) (interface{}, error) {

	db := mongo.GetDb()
	defer db.Session.Close()

	var contract entity.Contract
	contractColl := db.C(contract.CollectionName())

	// 将客户端提交用户信息保存
	reqData.Id = bson.NewObjectId()
	reqData.CustomerId = bson.NewObjectId()
	if err := contractColl.Insert(reqData); err != nil {
		return nil, err
	}

	return nil, nil
}

// 修改数据
func (this_ *Contract) Update(reqData *entity.Contract) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var contract entity.Contract
	contractColl := db.C(contract.CollectionName())

	// 对用户信息进行修改
	if err = contractColl.Update(bson.M{"_id": reqData.Id}, reqData); err != nil {
		return nil, err
	}

	return nil, err
}

// 删除数据
func (this_ *Contract) Delete(reqData *entity.Contract) (interface{}, error) {

	var err error

	db := mongo.GetDb()
	defer db.Session.Close()

	var contract entity.Contract
	contractColl := db.C(contract.CollectionName())

	if err = contractColl.RemoveId(reqData.Id); err != nil {
		return nil, err
	}

	return nil, err
}

// 添加文件名到files字段
func (this_ *Contract) AddFile(reqData *entity.Contract) (interface{}, error) {

	db := mongo.GetDb()
	defer db.Session.Close()

	var contract entity.Contract
	contractColl := db.C(contract.CollectionName())
	// 将客户端提交用户信息保存
	if err := contractColl.Update(bson.M{"_id": reqData.Id}, reqData); err != nil {
		return nil, err
	}

	return nil, nil
}

// 删除文件
func (this_ *Contract) DeleteFile(reqData *entity.Contract) (interface{}, error) {

	db := mongo.GetDb()
	defer db.Session.Close()

	var contract entity.Contract
	contractColl := db.C(contract.CollectionName())

	// 将客户端提交用户信息保存

	// if err = contractColl.RemoveId(reqData.Id); err != nil {
	// 	return nil, err
	// }
	if err := contractColl.Update(bson.M{"_id": reqData.Id}, bson.M{"$pull": bson.M{"files": reqData.Name}}); err != nil {
		return nil, err
	}

	fmt.Println("reqData.Files------------")
	fmt.Println(reqData.Files)

	return nil, nil
}
