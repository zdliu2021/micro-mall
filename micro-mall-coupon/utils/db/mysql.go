package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-demo/micro-mall-coupon/conf"
	"mall-demo/micro-mall-coupon/global"
	"mall-demo/micro-mall-coupon/model"
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
	global.SmsMysqlConn.AutoMigrate(&model.SmsCoupon{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsCouponHistory{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsCouponSpuRelation{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsHomeSubject{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsHomeAdv{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsCouponSpuCategoryRelation{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsHomeSubjectSpu{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsMemberPrice{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSeckillPromotion{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSeckillSession{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSeckillSkuNotice{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSeckillSkuRelation{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSkuFullReduction{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSkuLadder{})
	global.SmsMysqlConn.AutoMigrate(&model.SmsSpuBounds{})
}

func InitMysqlClient() *gorm.DB {
	global.GVA_CONFIG.CommonMysql.Dbname = "mall_sms"
	return GetMysqlInstance(global.GVA_CONFIG.CommonMysql)
}
