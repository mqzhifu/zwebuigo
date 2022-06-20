package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
)

type StatisticsLogService struct {
}

// CreateStatisticsLog 创建StatisticsLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (statisticsLogService *StatisticsLogService) CreateStatisticsLog(statisticsLog op.StatisticsLog) (err error) {

	err = global.GetGlobalDBByDBName("business").Create(&statisticsLog).Error
	return err
}

// DeleteStatisticsLog 删除StatisticsLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (statisticsLogService *StatisticsLogService)DeleteStatisticsLog(statisticsLog op.StatisticsLog) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&statisticsLog).Error
	return err
}

// DeleteStatisticsLogByIds 批量删除StatisticsLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (statisticsLogService *StatisticsLogService)DeleteStatisticsLogByIds(ids request.IdsReq) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&[]op.StatisticsLog{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStatisticsLog 更新StatisticsLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (statisticsLogService *StatisticsLogService)UpdateStatisticsLog(statisticsLog op.StatisticsLog) (err error) {
	err = global.GetGlobalDBByDBName("business").Save(&statisticsLog).Error
	return err
}

// GetStatisticsLog 根据id获取StatisticsLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (statisticsLogService *StatisticsLogService)GetStatisticsLog(id uint) (statisticsLog op.StatisticsLog, err error) {
	err = global.GetGlobalDBByDBName("business").Where("id = ?", id).First(&statisticsLog).Error
	return
}

// GetStatisticsLogInfoList 分页获取StatisticsLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (statisticsLogService *StatisticsLogService)GetStatisticsLogInfoList(info opReq.StatisticsLogSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GetGlobalDBByDBName("business").Model(&op.StatisticsLog{})
    var statisticsLogs []op.StatisticsLog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ProjectId != nil {
        db = db.Where("project_id = ?",info.ProjectId)
    }
    if info.Uid != nil {
        db = db.Where("uid = ?",info.Uid)
    }
    if info.Action != "" {
        db = db.Where("action LIKE ?","%"+ info.Action+"%")
    }
    if info.HeaderCommon != "" {
        db = db.Where("header_common LIKE ?","%"+ info.HeaderCommon+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&statisticsLogs).Error
	return  statisticsLogs, total, err
}
