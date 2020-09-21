package main

import (
	_ "echo_api/docs"
	//_ "github.com/go-openapi/swag"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"sync"
)

func main() {
	e := echo.New()
	/*e.GET("/", func(c echo.Context) error {
		//conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())

		//var conn *grpc.ClientConn
		// 建立对象
		var pipe = &sync.Pool{
			New: func() interface{} {
				conn, _ := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
				return conn
			},
		}

		conn := pipe.Get().(*grpc.ClientConn)

		defer conn.Close()
		client := pb.NewHelloServiceClient(conn)
		var params pb.HelloWorldRequest
		params.SkuId = []int32{1, 2, 3}
		result, err := client.HelloWorld(context.Background(), &params)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(500, err)
		}
		return c.JSON(200, result)
	})

	e.GET("/redis", func(c echo.Context) error {
		redis := ConnectClusterRedis()
		defer redis.Close()
		result := redis.HGetAll("stock:1000070001")
		return c.JSON(200, result)
	})

	e.GET("/test", func(c echo.Context) error {



		options := &pool.Options{
			InitTargets:  []string{"127.0.0.1:5000"},
			InitCap:      5,
			MaxCap:       30,
			DialTimeout:  time.Second * 5,
			IdleTimeout:  time.Second * 60,
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 5,
		}
		p, err := pool.NewGRPCPool(options, grpc.WithInsecure())
		if err != nil {
			fmt.Println("%#v\n", err)
		}
		if p == nil {
			fmt.Println("p= %#v\n", p)
		}
		defer p.Close()
		conn, err := p.Get()
		if err != nil {
			fmt.Println("%#v\n", err)
		}
		defer p.Put(conn)
		client := pb.NewHelloServiceClient(conn)
		var params pb.HelloWorldRequest
		params.SkuId = []int32{1, 2, 3}
		result, err := client.HelloWorld(context.Background(), &params)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(500, err)
		}
		return c.JSON(200, result)
	})*/
	//添加swagger文档地址
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/hello", Hello)

	e.Logger.Fatal(e.Start(":5001"))
}

//redis链接
func ConnectClusterRedis() *redis.Client {

	var pipe = &sync.Pool{
		New: func() interface{} {
			client := redis.NewClient(&redis.Options{
				Addr:     "10.1.1.245:6979",
				Password: "qlwhIPO82@#KDFQAwe",
				DB:       0,
			})
			return client
		},
	}
	conn := pipe.Get().(*redis.Client)
	//_, err := client.Ping().Result()
	//if err != nil {
	//	panic(err)
	//}
	return conn
}

// @Summary 门店内商品搜索
// @Tags 商品相关
// @Accept plain
// @Produce json
// @Param channel_id query int true "渠道id"
// @Param finance_code query string true "门店编码"
// @Param keywords query string false "搜索关键词"
// @Param page_index query int false "分页索引"
// @Param page_size query int false "分页大小"
// @Success 200 {object} BaseResponse
// @Failure 400 {object} BaseResponse
// @Router /product-api/product/search [get]
func Hello(context echo.Context) error {
	return nil
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
