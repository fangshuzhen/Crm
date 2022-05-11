package micro

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/core/config"
	"hello/core/utils"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

type ClientDis struct {
	client     *clientv3.Client
	ServerList map[string]Server
	lock       sync.Mutex
}

//微服务客户端
var microClient *ClientDis

var etcdService *ServiceReg

func GetNewClientDis(config config.MicroConfig) (*ClientDis, error) {
	if microClient == nil {
		conf := clientv3.Config{
			Endpoints:   config.Etcd,
			DialTimeout: 5 * time.Second,
		}

		if client, err := clientv3.New(conf); err != nil {
			return nil, err
		} else {
			microClient = &ClientDis{
				client:     client,
				ServerList: make(map[string]Server),
			}
		}
	}
	return microClient, nil
}

//提取地址
func (this_ *ClientDis) extractAddress(resp *clientv3.GetResponse, pub func(putKey, putValue string)) []string {

	address := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return address
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			this_.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value), pub)
			address = append(address, string(v))
		}
	}
	return address
}

//添加服务
func (this_ *ClientDis) SetServiceList(key, val string, pub func(putKey, pubValue string)) {
	this_.lock.Lock()
	defer this_.lock.Unlock()

	var server Server
	json.Unmarshal([]byte(val), &server)

	this_.ServerList[key] = server

	if pub != nil {
		pub(key, val)
	}

	log.Println("set data : ", key, "value:", val)
}

//删除服务
func (this_ *ClientDis) DelServiceList(key string, del func(key string)) {
	this_.lock.Lock()
	defer this_.lock.Unlock()
	delete(this_.ServerList, key)
	if del != nil {
		del(key)
	}
	log.Println("delete data key:", key)
}

//监听ETCD的操作
func (this_ *ClientDis) watcher(prefix string, put func(key, value string), del func(key string)) {
	rch := this_.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type.String() {
			case "PUT":
				this_.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value), put)
			case "DELETE":
				this_.DelServiceList(string(ev.Kv.Key), del)
			}
		}
	}
}

//获得服务列表
func (this_ *ClientDis) GetService(prefix string, put func(key, value string), del func(key string)) ([]string, error) {
	resp, err := this_.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := this_.extractAddress(resp, put)

	go this_.watcher(prefix, put, del)

	return addrs, nil
}

//筛选服务
func (this_ *ClientDis) GetMicro(name string) (*Server, error) {
	candidateServer := make([]Server, 0)
	fmt.Println(this_.ServerList)
	for _, server := range this_.ServerList {
		if server.Name == name {
			candidateServer = append(candidateServer, server)
		}
	}

	//随机筛选一个服务
	if len(candidateServer) > 0 {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(len(candidateServer))
		return &candidateServer[n], nil
	}

	return nil, fmt.Errorf("没有找到服务:%s", name)
}

//监听服务上线与上线
func (this_ *ClientDis) Listen() {
	this_.GetService("/micro-server/registry", func(key, value string) {
		fmt.Println("添加服务:", key, value)
	}, func(key string) {
		fmt.Println("删除服务:", key)
	})
}

//从ETCD下线
func UnregisterEtcd() {
	etcdService.RevokeLease()
}

//将服务注册到ETCD
func RegisterToEtcd(conf config.MicroConfig) {
	var err error
	etcdService, err = NewServiceReg(conf.Etcd, 20)
	if err == nil {
		var server = Server{
			Name:    conf.Name,
			Address: conf.Advertise,
		}
		key := fmt.Sprintf("/micro-server/registry/%s-%s", conf.Name, utils.GetUuid())
		etcdService.PutServer(key, server)
	} else {
		panic(err)
	}
}

//监听ectd
func ListenEtcd(conf config.MicroConfig) {
	log.Printf("监听服务:%s\n", conf.Address)
	microClient, err := GetNewClientDis(conf)
	if err != nil {
		fmt.Printf("启动etcd客户端监听失败:%s", err)
	}
	microClient.Listen()
}

//启动gRPC，注册到ETCD，开启监听
func RunGrpc(conf config.MicroConfig, registerGreeter func(server *grpc.Server)) {
	ListenEtcd(conf)

	//监听本地端口
	lis, err := net.Listen("tcp", conf.Address)
	if err != nil {
		fmt.Printf("监听端口失败:%s", err)
		return
	}

	//创建gRpc服务
	server := grpc.NewServer()
	registerGreeter(server)

	//注册
	RegisterToEtcd(conf)

	if err = server.Serve(lis); err != nil {
		fmt.Printf("开启服务失败: %s", err)
	} else {
		fmt.Printf("开启服务成功")
	}

}
