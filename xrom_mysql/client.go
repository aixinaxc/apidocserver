package xrom_mysql

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"apidocserver/base"
)

func Client() *xorm.Engine {
	engine,err := xorm.NewEngine("mysql",base.MysqlUrl)
	if err != nil {
		fmt.Println("mysql打开失败",err)
	}
	engine.SetConnMaxLifetime(base.MysqlMaxLifetime)
	engine.SetMaxOpenConns(base.MysqlMaxOpenConn)
	engine.SetMaxIdleConns(base.MysqlMaxIdleConn)
	engine.ShowSQL(true)
	err = engine.Ping()
	if err != nil {
		fmt.Println("mysql没有ping成功",err)
		return nil
	}
	return engine
}

func InsertXORMMsg(t interface{}) bool {
	engine := Client()
	if engine == nil {
		fmt.Println("mysql连接失败")
		return false
	}

	_,err := engine.Insert(t)
	if err != nil {
		fmt.Println("数据插入失败",err)
		return false
	}
	return true
}

func FindXROMMsg(sql string) []map[string][]byte  {
	engine := Client()
	if engine == nil {
		fmt.Println("mysql连接失败")
		return nil
	}
	results,err := engine.Query(sql)
	if err != nil {
		fmt.Println("查询数据失败")
		return nil
	}
	return results
}

