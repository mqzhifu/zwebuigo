// 自动生成模板Project
package op

import "github.com/flipped-aurora/gin-vue-admin/server/model"

// Project 结构体
type Project struct {
	//global.GVA_MODEL
	model.DIY_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	Type      *int   `json:"type" form:"type" gorm:"column:type;comment:类型,1service 2frontend 3backend 4app;"`
	Access    string `json:"access" form:"access" gorm:"column:access;comment:baseAuth 认证KEY;size:255;"`
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述信息;size:255;"`
	Git       string `json:"git" form:"git" gorm:"column:git;comment:git仓地址;size:255;"`
	Lang      *int   `json:"lang" form:"lang" gorm:"column:lang;comment:实现语言1php2go3java4js;"`
	SecretKey string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:密钥;size:100;"`
	Status    *int   `json:"status" form:"status" gorm:"column:status;comment:状态1正常2关闭;"`
}

// TableName Project 表名
func (Project) TableName() string {
	return "project"
}
