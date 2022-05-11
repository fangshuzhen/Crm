package micro_controller

import (
	"context"
	"fmt"
	"hello/core/proto/pb"
	"hello/user-center/internal/service"
)

type User struct{}

func (this_ *User) GetUserList(c context.Context, reqData *pb.RequestUserList) (*pb.ReplyUserList, error) {
	fmt.Println("我正大查询数据库....")

	replay := pb.ReplyUserList{}

	var userList []*pb.User

	userList = append(userList, &pb.User{Id: "1", Username: "小明同学", Name: "xxx"})
	userList = append(userList, &pb.User{Id: "2", Username: "小明同学2", Name: "xxx"})

	replay.UserList = userList

	return &replay, nil
}

func (this_ *User) CheckToken(c context.Context, reqData *pb.RequestCheckToken) (*pb.User, error) {

	userInfo, err := service.GetAdminService().Check(reqData.Token)
	if err != nil {
		return nil, err
	}

	//封装grpc的返回信息
	replyUser := pb.User{
		Id:           userInfo.Id.Hex(),
		Username:     userInfo.Username,
		Name:         userInfo.Name,
		Avatar:       userInfo.Avatar,
		Introduction: userInfo.Introduction,
		Roles:        userInfo.Roles,
		Token:        userInfo.Token,
	}

	return &replyUser, nil
}
