package mongo

import (
	"hello/core/config"
	"log"
	"time"

	"github.com/globalsign/mgo"
	_ "github.com/globalsign/mgo/bson"
)

var session *mgo.Session

var mongoConf *config.MongoDb

func DBConnect(conf *config.MongoDb) error {

	mongoConf = conf

	info := &mgo.DialInfo{
		Addrs:          mongoConf.Url,
		Timeout:        time.Duration(mongoConf.Timeout) * time.Second,
		Database:       mongoConf.Name,
		Username:       mongoConf.UserName,
		Password:       mongoConf.Password,
		ReplicaSetName: mongoConf.ReplicaSetName,
		Source:         mongoConf.Source,
	}

	var err error

	session, err = mgo.DialWithInfo(info)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	session.SetSocketTimeout(time.Duration(mongoConf.Timeout) * time.Second)

	//设置连接池的大小
	if mongoConf.MaxPoolSize > 0 {
		session.SetPoolLimit(mongoConf.MaxPoolSize)
	}

	log.Printf("MongoDB连接成功！")

	return nil
}

func GetDb() *mgo.Database {
	return session.Clone().DB(mongoConf.Name)
}
