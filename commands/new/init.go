//package commands
package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"

	//"github.com/gomodule/redigo/redis"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct{
	Id 	int
	Name string
}

//var mysqlInit = `package models (表是脚手架运行后生成的包)

//import (
//		"context"
//		"errors"
//		"fmt"
//		//_ "log"
//		"time"
//		"github.com/gomodule/redigo/redis"
//		"gorm.io/driver/mysql"
//		//orm "gorm.io/gorm"
//		//_ "github.com/gomodule/redigo/redis"
//		"github.com/spf13/viper"
//		"github.com/qiniu/qmgo"
//		//jinzhu "github.com/jinzhu/gorm"
//)

var (
	//Jinzhu *jinzhu.DB
	Mysql *gorm.DB
	//Redis *redis.Conn
	//Cli *qmgo.QmgoClient
	err error
)


func init(){
	// 1. connection mysql database
	 initMysql()
	 initRedis()
	//Jinzhu,err = initJinzhu()
	// 2.自动迁移(自动建表)
	// gorm@v1.20.2
	//Mysql.AutoMigrate(




		//)

	// 3. gorm@v1.19.X





		//)
}

// ======== ======== ======== ======== ======== ======== ======== ======== ======== ========
// 初始化mysql
func initMysql() (*gorm.DB,error){
	Mysql, err = gorm.Open(mysql.Open(GetMysqlConf()), &gorm.Config{})
	if err != nil {
		err = errors.New("mysql connect failed")
	}
	return Mysql,err
}

// 初始化mysql数据库连接
//func initJinzhu()(*jinzhu.DB,error){
//	Jinzhu, err = jinzhu.Open("mysql", GetMysqlConf())
//	defer func() {
//		if err = Jinzhu.Close();err != nil{
//			panic(err)
//		}
//	}()
//	Jinzhu.SingularTable(true)
//	return Jinzhu,err
//}

// 获取配置文件中 mysql的参数
func GetMysqlConf() string{
	return MysqlConfig(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.database"),
	)
}

// 拼接mysql参数
func MysqlConfig(username,password,addr,database string) string{
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		database,
		true,
		"Local")
	return config
}
// ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ========

// 初始化redis
func initRedis() (*redis.Conn,error){
	pool := &redis.Pool{
		MaxActive: viper.Get("maxActive").(int),
		MaxIdle: viper.Get("maxIdle").(int),
		Wait: viper.Get("wait").(bool),
		IdleTimeout: time.Duration(viper.Get("count").(int64)) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",viper.GetString("addr"),redis.DialDatabase(viper.GetInt("db")))
		},
	}
	conn := pool.Get()

	if r,_ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed")
	}

	defer func(){
		if err = conn.Close();err != nil {
			panic(err)
		}
	}()
	return &conn,err
}
// ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ========
// MongoDB 客户端连接实例
func initMongoDB(){
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: viper.GetString("uri"),
		Database: viper.GetString("database"),
		Coll:     viper.GetString("coll")})
	defer func(){
		if err = cli.Close(ctx);err != nil {
			panic(err)
		}
	}()
}
// ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ======== ========
























