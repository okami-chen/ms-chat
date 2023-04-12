package initialize

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/okamin-chen/chat/pkg/global"
)

// 初始化时如果没有表创建表
func InitTable() {
	m := global.Conf.Server
	if m.IsInitTable == true {
		biz.ErrIsNil(
			global.Db.AutoMigrate(),
			"初始化表失败")
	}
}
