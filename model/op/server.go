// 自动生成模板Server
package op

import "github.com/flipped-aurora/gin-vue-admin/server/model"

// Server 结构体
type Server struct {
	//global.GVA_MODEL
	model.DIY_MODEL
	Name           string `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	Env            *int   `json:"env" form:"env" gorm:"column:env;comment:环境变量,1本地2开发3测试4预发布5线上;size:10;"`
	InnerIp        string `json:"innerIp" form:"innerIp" gorm:"column:inner_ip;comment:内网IP;size:15;"`
	OutIp          string `json:"outIp" form:"outIp" gorm:"column:out_ip;comment:外网IP;size:15;"`
	Platform       *int   `json:"platform" form:"platform" gorm:"column:platform;comment:平台类型1自有2阿里3腾讯4华为;size:10;"`
	ChargeUserName string `json:"chargeUserName" form:"chargeUserName" gorm:"column:charge_user_name;comment:负责人姓名;size:50;"`
	Price          *int   `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`
	Ext            string `json:"ext" form:"ext" gorm:"column:ext;comment:自定义配置信息;size:255;"`
	EndTime        *int   `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间;size:10;"`
	StartTime      *int   `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间;size:10;"`
	Status         *int   `json:"status" form:"status" gorm:"column:status;comment:状态1正常2关闭;"`
}

// TableName Server 表名
func (Server) TableName() string {
	return "server"
}
