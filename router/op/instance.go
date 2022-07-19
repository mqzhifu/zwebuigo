package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InstanceRouter struct {
}

// InitInstanceRouter 初始化 Instance 路由信息
func (s *InstanceRouter) InitInstanceRouter(Router *gin.RouterGroup) {
	instanceRouter := Router.Group("instance").Use(middleware.OperationRecord())
	instanceRouterWithoutRecord := Router.Group("instance")
	var instanceApi = v1.ApiGroupApp.OpApiGroup.InstanceApi
	{
		instanceRouter.POST("createInstance", instanceApi.CreateInstance)   // 新建Instance
		instanceRouter.DELETE("deleteInstance", instanceApi.DeleteInstance) // 删除Instance
		instanceRouter.DELETE("deleteInstanceByIds", instanceApi.DeleteInstanceByIds) // 批量删除Instance
		instanceRouter.PUT("updateInstance", instanceApi.UpdateInstance)    // 更新Instance
	}
	{
		instanceRouterWithoutRecord.GET("findInstance", instanceApi.FindInstance)        // 根据ID获取Instance
		instanceRouterWithoutRecord.GET("getInstanceList", instanceApi.GetInstanceList)  // 获取Instance列表
	}
}
