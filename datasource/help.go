package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	//"sync"
	"go-trial/config"
	"log"
)

var (
	ormEngine *xorm.Engine
	//lock      sync.Locker
)

func Instance() *xorm.Engine {

	if ormEngine != nil {
		return ormEngine
	}
	//lock.Lock()
	//defer lock.Unlock()

	if ormEngine != nil {
		return ormEngine
	}
	db := config.Db
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		db.User, db.Pass, db.Host, db.Port, db.Name)
	engine, err := xorm.NewEngine(config.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstance,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	engine.SetTZLocation(config.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cache)
	ormEngine = engine
	return engine
}

