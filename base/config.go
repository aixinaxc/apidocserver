package base

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)

var RedisUrl string
var RedisPassword string
var MysqlUrl string


func Config()  {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println("config:",err)
		return
	}
	RedisUrl,err = config.Get("RedisUrl")
	RedisPassword,err = config.Get("RedisPassword")
	MysqlUrl,err = config.Get("MysqlUrl")
	if err != nil{
		fmt.Println("config-data:",err)
	}
	fmt.Println("redis_config:",RedisUrl,RedisPassword)
	fmt.Println("mysql_config:",MysqlUrl)
}
