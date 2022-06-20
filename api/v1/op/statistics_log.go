package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/op"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type StatisticsLogApi struct {
}

var statisticsLogService = service.ServiceGroupApp.OpServiceGroup.StatisticsLogService


// CreateStatisticsLog 创建StatisticsLog
// @Tags StatisticsLog
// @Summary 创建StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.StatisticsLog true "创建StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /statisticsLog/createStatisticsLog [post]
func (statisticsLogApi *StatisticsLogApi) CreateStatisticsLog(c *gin.Context) {
	var statisticsLog op.StatisticsLog
	_ = c.ShouldBindJSON(&statisticsLog)
	if err := statisticsLogService.CreateStatisticsLog(statisticsLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStatisticsLog 删除StatisticsLog
// @Tags StatisticsLog
// @Summary 删除StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.StatisticsLog true "删除StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /statisticsLog/deleteStatisticsLog [delete]
func (statisticsLogApi *StatisticsLogApi) DeleteStatisticsLog(c *gin.Context) {
	var statisticsLog op.StatisticsLog
	_ = c.ShouldBindJSON(&statisticsLog)
	if err := statisticsLogService.DeleteStatisticsLog(statisticsLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStatisticsLogByIds 批量删除StatisticsLog
// @Tags StatisticsLog
// @Summary 批量删除StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /statisticsLog/deleteStatisticsLogByIds [delete]
func (statisticsLogApi *StatisticsLogApi) DeleteStatisticsLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := statisticsLogService.DeleteStatisticsLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStatisticsLog 更新StatisticsLog
// @Tags StatisticsLog
// @Summary 更新StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.StatisticsLog true "更新StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /statisticsLog/updateStatisticsLog [put]
func (statisticsLogApi *StatisticsLogApi) UpdateStatisticsLog(c *gin.Context) {
	var statisticsLog op.StatisticsLog
	_ = c.ShouldBindJSON(&statisticsLog)
	if err := statisticsLogService.UpdateStatisticsLog(statisticsLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStatisticsLog 用id查询StatisticsLog
// @Tags StatisticsLog
// @Summary 用id查询StatisticsLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query op.StatisticsLog true "用id查询StatisticsLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /statisticsLog/findStatisticsLog [get]
func (statisticsLogApi *StatisticsLogApi) FindStatisticsLog(c *gin.Context) {
	var statisticsLog op.StatisticsLog
	_ = c.ShouldBindQuery(&statisticsLog)
	if restatisticsLog, err := statisticsLogService.GetStatisticsLog(statisticsLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restatisticsLog": restatisticsLog}, c)
	}
}

// GetStatisticsLogList 分页获取StatisticsLog列表
// @Tags StatisticsLog
// @Summary 分页获取StatisticsLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query opReq.StatisticsLogSearch true "分页获取StatisticsLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /statisticsLog/getStatisticsLogList [get]
func (statisticsLogApi *StatisticsLogApi) GetStatisticsLogList(c *gin.Context) {
	var pageInfo opReq.StatisticsLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := statisticsLogService.GetStatisticsLogInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
