package starter

import (
	"github.com/XM-GO/PandaKit/logger"
	"github.com/okamin-chen/chat/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	_ "github.com/lib/pq"
)

var Db *gorm.DB

type DbGorm struct {
	Type         string
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
}

func (dg *DbGorm) GormInit() *gorm.DB {
	dg.Dsn = global.Conf.Mysql.Dsn()
	dg.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
	dg.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
	switch dg.Type {
	default:
		Db = dg.GormMysql()
	}
	return Db
}
func (dg *DbGorm) GormMysql() *gorm.DB {

	mysqlConfig := mysql.Config{
		DSN:                       dg.Dsn, // DSN data source name
		DefaultStringSize:         191,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,  // 根据版本自动配置
	}
	ormConfig := &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), ormConfig)
	if err != nil {
		logger.Log.Panicf("连接mysql失败! [%s]", err.Error())
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dg.MaxOpenConns)
	global.Log.Infoln("Mysql Init Success")
	return db
}
