// 自动生成模板StatisticsLog
package op

import "github.com/flipped-aurora/gin-vue-admin/server/model"

// StatisticsLog 结构体
type StatisticsLog struct {
	//global.GVA_MODEL
	model.DIY_MODEL
	ProjectId     *int   `json:"projectId" form:"projectId" gorm:"column:project_id;comment:项目ID;size:10;"`
	Uid           *int   `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;size:10;"`
	Sn            string `json:"sn" form:"sn" gorm:"column:sn;comment:设备-序列号;size:100;"`
	Action        string `json:"action" form:"action" gorm:"column:action;comment:动作标识;size:255;"`
	AppName       string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名;size:100;"`
	AppVersion    string `json:"appVersion" form:"appVersion" gorm:"column:app_version;comment:应用版本号;size:100;"`
	Msg           string `json:"msg" form:"msg" gorm:"column:msg;comment:自定义消息体;size:255;"`
	PackageName   string `json:"packageName" form:"packageName" gorm:"column:package_name;comment:包名;size:100;"`
	RecordTime    *int   `json:"recordTime" form:"recordTime" gorm:"column:record_time;comment:记录时间;size:10;"`
	SystemVersion string `json:"systemVersion" form:"systemVersion" gorm:"column:system_version;comment:设备-版本号;size:100;"`
	HeaderCommon  string `json:"headerCommon" form:"headerCommon" gorm:"column:header_common;comment:http公共请求头信息;"`
	HeaderBase    string `json:"headerBase" form:"headerBase" gorm:"column:header_base;comment:http请求头客户端基础信息;"`
}

// TableName StatisticsLog 表名
func (StatisticsLog) TableName() string {
	return "statistics_log"
}
