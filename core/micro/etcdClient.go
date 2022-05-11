package micro

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type ServiceReg struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	cancelfun     func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse //channel
	key           string
}

type Server struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewServiceReg(addr []string, timeNum int64) (*ServiceReg, error) {

	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}

	var err error
	var client *clientv3.Client

	//实例化一个客户端
	if client, err = clientv3.New(conf); err != nil {
		return nil, err
	}
	//实例化一个租约服务
	server := &ServiceReg{
		client: client,
	}
	if err := server.setLease(timeNum); err != nil {
		return nil, err
	}
	go server.ListenLeaseRespChan()
	return server, nil
}

//设置续租
func (this_ *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(this_.client)
	//设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), timeNum)
	if err != nil {
		return err
	}
	//设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}
	this_.lease = lease
	this_.leaseResp = leaseResp
	this_.cancelfun = cancelFunc
	this_.keepAliveChan = leaseRespChan
	return nil
}

//监听续租情况
func (this_ *ServiceReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-this_.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\nW")
				return
			} else {
				fmt.Printf("续租成功\n")
			}
		}
	}
}

//GetValue 从ETCD中读取数据
func (this_ *ServiceReg) GetValue(prefix string) ([]string, error) {
	//指定前缀到ETCD中去读取服务列表
	resp, err := this_.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := make([]string, 0)
	//如果读取不到数据，则返回一个长度为0的数组
	if resp == nil || resp.Kvs == nil {
		return addrs, nil
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			addrs = append(addrs, string(v))
		}
	}
	return addrs, nil
}

//PutServer 读取ETCD写入微服务的配置
func (this_ *ServiceReg) PutServer(key string, server Server) error {
	kv := clientv3.NewKV(this_.client)
	cTodo := context.TODO()
	by, _ := json.Marshal(server)
	_, err := kv.Put(cTodo, key, string(by))
	return err
}

//PutValue 读取ETCD写入数据
func (this_ *ServiceReg) PutValue(key string, value interface{}) error {
	kv := clientv3.NewKV(this_.client)
	cTodo := context.TODO()
	by, _ := json.Marshal(value)
	_, err := kv.Put(cTodo, key, string(by))
	return err
}

//RevokeLease 撤销续租
func (this_ *ServiceReg) RevokeLease() error {
	this_.cancelfun()
	time.Sleep(2 * time.Second)
	_, err := this_.lease.Revoke(context.TODO(), this_.leaseResp.ID)
	return err
}
