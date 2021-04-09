package datas

import (
	"gin-vue-devops/model"
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

func InitMysqlData(db *gorm.DB) {
	InitSysApi(db)
	InitSysUser(db)
	InitCasbinModel(db)
	InitSysAuthority(db)
	InitSysBaseMenus(db)
	InitAuthorityMenu(db)
<<<<<<< HEAD
	InitSysDictionary(db)
	InitSysAuthorityMenus(db)
	InitSysDataAuthorityId(db)
	InitSysDictionaryDetail(db)
=======
	InitSysAuthorityMenus(db)
	InitSysDataAuthorityId(db)
>>>>>>> develop
	InitExaFileUploadAndDownload(db)
}

func InitMysqlTables(db *gorm.DB) {
	var err error
	if !db.Migrator().HasTable("casbin_rule") {
		err = db.Migrator().CreateTable(&gormadapter.CasbinRule{})
	}
	err = db.AutoMigrate(
		model.SysApi{},
		model.SysUser{},
		model.SysBaseMenu{},
		model.SysAuthority{},
		model.JwtBlacklist{},
<<<<<<< HEAD
		model.SysDictionary{},
		model.SysDictionaryDetail{},
=======
>>>>>>> develop
		model.SysBaseMenuParameter{},
		model.ExaFileUploadAndDownload{},
	)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql]-->初始化数据表成功")
}
