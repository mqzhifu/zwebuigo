package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StatisticsLogRouter struct {
}

// InitStatisticsLogRouter 初始化 StatisticsLog 路由信息
func (s *StatisticsLogRouter) InitStatisticsLogRouter(Router *gin.RouterGroup) {
	statisticsLogRouter := Router.Group("statisticsLog").Use(middleware.OperationRecord())
	statisticsLogRouterWithoutRecord := Router.Group("statisticsLog")
	var statisticsLogApi = v1.ApiGroupApp.OpApiGroup.StatisticsLogApi
	{
		statisticsLogRouter.POST("createStatisticsLog", statisticsLogApi.CreateStatisticsLog)   // 新建StatisticsLog
		statisticsLogRouter.DELETE("deleteStatisticsLog", statisticsLogApi.DeleteStatisticsLog) // 删除StatisticsLog
		statisticsLogRouter.DELETE("deleteStatisticsLogByIds", statisticsLogApi.DeleteStatisticsLogByIds) // 批量删除StatisticsLog
		statisticsLogRouter.PUT("updateStatisticsLog", statisticsLogApi.UpdateStatisticsLog)    // 更新StatisticsLog
	}
	{
		statisticsLogRouterWithoutRecord.GET("findStatisticsLog", statisticsLogApi.FindStatisticsLog)        // 根据ID获取StatisticsLog
		statisticsLogRouterWithoutRecord.GET("getStatisticsLogList", statisticsLogApi.GetStatisticsLogList)  // 获取StatisticsLog列表
	}
}
