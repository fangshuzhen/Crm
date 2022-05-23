package main

import (
	"context"
	"flag"
	"fmt"
	"hello/core/config"
	myhttp "hello/core/http"
	"hello/core/micro"
	"hello/core/middleware"
	"hello/core/mongo"
	"hello/log-center/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {

	res := flag.String("res", "./res/configuration.toml", "Res:配置文件")
	flag.Parse()

	err := config.InitConfig(*res)
	if err != nil {
		log.Println("读取配置失败:", err)
		panic(err)
	} else {
		log.Println("读取配置成功:", config.Conf)
	}

	log.Println("开始连接MongoDB")
	if err = mongo.DBConnect(&config.Conf.Mongo); err != nil {
		log.Println("MongDB连接失败")
		panic(err)
	} else {
		log.Println("MongDB连接成功")
	}
}

func main() {

	// 启动微服务
	go micro.RunGrpc(config.Conf.Micro, func(server *grpc.Server) {
		// TODO ：注册微服务方法
		// pb.RegisterGreeterUserServer(server, &micro_controller.User{})

		reflection.Register(server)
	})

	// 只有监听了ETCD才能拿到微服务
	go micro.ListenEtcd(config.Conf.Micro)

	router := gin.Default()
	router.Use(myhttp.Cors())
	route.LoadRoute(router)

	//日志记录
	// router.Use(middleware.LoggerToFile())
	router.Use(middleware.LoggerToMongo())

	s := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.Conf.Server.Host, config.Conf.Server.Port),
		Handler:           router,
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	// h2s := &http2.Server{IdleTimeout: 1 * time.Minute, PermitProhibitedCipherSuites: true}
	// http2.ConfigureServer(s, h2s)

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)
		<-quit
		log.Panicln("Shutting down server......")

		// 服务关闭时从ETCD中移除注册（只有注册了服务，才有这个必要）
		micro.UnregisterEtcd()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")

	}()
	// log.Printf("监听http服务：%s:%d", config.Conf.Server.Host, config.Conf.Server.Port)

	// 启动http监听
	s.ListenAndServe()

}
