package service

import (
	"context"
	"hello/core/config"
	"hello/core/entity"
	"hello/core/micro"
	"hello/core/proto/pb"

	"github.com/globalsign/mgo/bson"
	"google.golang.org/grpc"
)

var microUserConn *grpc.ClientConn

func getUserConn() (*grpc.ClientConn, error) {
	if microUserConn == nil {
		var client *micro.ClientDis
		var server *micro.Server
		var err error

		if client, err = micro.GetNewClientDis(config.Conf.Micro); err != nil {
			return nil, err
		}

		if server, err = client.GetMicro("com.ipk.hello.userCenter"); err != nil {
			return nil, err
		}
		if microUserConn, err = grpc.Dial(server.Address, grpc.WithInsecure()); err != nil {
			return nil, err
		}
	}
	return microUserConn, nil
}

func CheckToken(token string) (*entity.UserInfo, error) {

	// 获得一个连接
	conn, err := getUserConn()
	if err != nil {
		return nil, err
	}
	// 获得一个客户端
	userClient := pb.NewGreeterUserClient(conn)
	// 构造请求参数
	reqData := pb.RequestCheckToken{Token: token}
	// 请求微服务
	var pbUser *pb.User
	if pbUser, err = userClient.CheckToken(context.TODO(), &reqData); err != nil {
		return nil, err
	}
	// 封装用户信息
	userInfo := entity.UserInfo{
		Id:           bson.ObjectIdHex(pbUser.Id),
		Username:     pbUser.Username,
		Name:         pbUser.Name,
		Avatar:       pbUser.Avatar,
		Introduction: pbUser.Introduction,
		Roles:        pbUser.Roles,
		Token:        pbUser.Token,
	}

	return &userInfo, nil
}
