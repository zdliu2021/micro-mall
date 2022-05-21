package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-demo/micro-mall-ware/conf"
	"mall-demo/micro-mall-ware/global"
	"mall-demo/micro-mall-ware/model"
)

func GetMysqlInstance(mysqlOption conf.MysqlOption) *gorm.DB {
	if db, err := gorm.Open(mysql.Open(mysqlOption.Dsn()), &gorm.Config{
		SkipDefaultTransaction: false,
	}); err != nil {
		fmt.Println("error" + err.Error())
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mysqlOption.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysqlOption.MaxOpenConns)
		return db
	}
}

func InitMysqlTables() {
	global.WmsMysqlConn.AutoMigrate(&model.WmsPurchase{})
	global.WmsMysqlConn.AutoMigrate(&model.UndoLog{})
	global.WmsMysqlConn.AutoMigrate(&model.WmsWareInfo{})
	global.WmsMysqlConn.AutoMigrate(&model.WmsPurchaseDetail{})
	global.WmsMysqlConn.AutoMigrate(&model.WmsWareSku{})
	global.WmsMysqlConn.AutoMigrate(&model.WmsWareOrderTask{})
	global.WmsMysqlConn.AutoMigrate(&model.WmsWareOrderTaskDetail{})
}

func InitMysqlClient() *gorm.DB {
	global.GVA_CONFIG.CommonMysql.Dbname = "mall_wms"
	return GetMysqlInstance(global.GVA_CONFIG.CommonMysql)
}
