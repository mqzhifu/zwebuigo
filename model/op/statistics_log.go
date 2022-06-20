// 自动生成模板StatisticsLog
package op

import "github.com/flipped-aurora/gin-vue-admin/server/model"

// StatisticsLog 结构体
// 如果含有time.Time 请自行import time包
type StatisticsLog struct {
      //global.GVA_MODEL
      model.DIY_MODEL
      ProjectId  *int `json:"projectId" form:"projectId" gorm:"column:project_id;comment:项目ID;size:10;"`
      Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;size:10;"`
      Category  *int `json:"category" form:"category" gorm:"column:category;comment:分类，暂未使用;size:10;"`
      Action  string `json:"action" form:"action" gorm:"column:action;comment:动作标识;size:255;"`
      HeaderBase  string `json:"headerBase" form:"headerBase" gorm:"column:header_base;comment:http请求头客户端基础信息;size:255;"`
      HeaderCommon  string `json:"headerCommon" form:"headerCommon" gorm:"column:header_common;comment:http公共请求头信息;"`
      Msg  string `json:"msg" form:"msg" gorm:"column:msg;comment:自定义消息体;size:255;"`
}


// TableName StatisticsLog 表名
func (StatisticsLog) TableName() string {
  return "statistics_log"
}

