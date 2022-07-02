package Tools

import (
	"blog/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/wonderivan/logger"
)

//type Orm struct {
//	*xorm.Engine
//}
var DbEngine *xorm.Engine

func OrmEngine(database model.DatabaseConfig) error {
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		logger.Error("xorm Engine 启动失败", err)
		return err
	}
	engine.ShowSQL(database.ShowSql)
	err = engine.Sync2(new(model.User)) //将结构体映射为数据库中的表
	if err != nil {
		logger.Error(err)
	}
	if err := engine.Ping(); err != nil { //用于测试数据库是否连接
		fmt.Println(err)
	}
	DbEngine = engine
	return nil
}
