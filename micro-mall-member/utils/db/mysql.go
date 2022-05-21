package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-demo/micro-mall-member/conf"
	"mall-demo/micro-mall-member/global"
	"mall-demo/micro-mall-member/model"
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
	global.UmsMysqlConn.AutoMigrate(&model.UmsMember{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsGrowthChangeHistory{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsMemberLevel{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsMemberCollectSpu{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsIntegrationChangeHistory{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsMemberLoginLog{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsMemberReceiveAddress{})
	global.UmsMysqlConn.AutoMigrate(&model.UmsMemberStatisticsInfo{})
}

func InitMysqlClient() (*gorm.DB) {
	global.GVA_CONFIG.CommonMysql.Dbname = "mall_ums"
	return GetMysqlInstance(global.GVA_CONFIG.CommonMysql)

}
