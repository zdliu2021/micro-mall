package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-demo/micro-mall-order/conf"
	"mall-demo/micro-mall-order/global"
	"mall-demo/micro-mall-order/model"
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
	global.OmsMysqlConn.AutoMigrate(&model.OmsOrder{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsOrderItem{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsOrderOperateHistory{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsOrderSetting{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsOrderReturnApply{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsOrderReturnReason{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsPaymentInfo{})
	global.OmsMysqlConn.AutoMigrate(&model.OmsRefundInfo{})
}

func InitMysqlClient() *gorm.DB {
	global.GVA_CONFIG.CommonMysql.Dbname = "mall_oms"
	return GetMysqlInstance(global.GVA_CONFIG.CommonMysql)
}
