package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-demo/micro-mall-product/conf"
	"mall-demo/micro-mall-product/global"
	"mall-demo/micro-mall-product/model"
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
	global.PmsMysqlConn.AutoMigrate(&model.PmsAttr{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsBrand{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsCategory{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsAttrAttrgroupRelation{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsAttrGroup{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsCommentReplay{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsProductAttrValue{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsSkuImages{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsSkuInfo{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsSkuSaleAttrValue{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsSpuInfoDesc{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsSpuImages{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsSpuComment{})
	global.PmsMysqlConn.AutoMigrate(&model.PmsCategoryBrandRelation{})
}

func InitMysqlClient() *gorm.DB {
	global.GVA_CONFIG.CommonMysql.Dbname = "mall_pms"
	return GetMysqlInstance(global.GVA_CONFIG.CommonMysql)

}
