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
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type InstanceApi struct {
}

var instanceService = service.ServiceGroupApp.OpServiceGroup.InstanceService


// CreateInstance 创建Instance
// @Tags Instance
// @Summary 创建Instance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.Instance true "创建Instance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /instance/createInstance [post]
func (instanceApi *InstanceApi) CreateInstance(c *gin.Context) {
	var instance op.Instance
	_ = c.ShouldBindJSON(&instance)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Platform":{utils.NotEmpty()},
        "Env":{utils.NotEmpty()},
        "Host":{utils.NotEmpty()},
        "User":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(instance, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := instanceService.CreateInstance(instance); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInstance 删除Instance
// @Tags Instance
// @Summary 删除Instance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.Instance true "删除Instance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /instance/deleteInstance [delete]
func (instanceApi *InstanceApi) DeleteInstance(c *gin.Context) {
	var instance op.Instance
	_ = c.ShouldBindJSON(&instance)
	if err := instanceService.DeleteInstance(instance); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInstanceByIds 批量删除Instance
// @Tags Instance
// @Summary 批量删除Instance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Instance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /instance/deleteInstanceByIds [delete]
func (instanceApi *InstanceApi) DeleteInstanceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := instanceService.DeleteInstanceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInstance 更新Instance
// @Tags Instance
// @Summary 更新Instance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body op.Instance true "更新Instance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /instance/updateInstance [put]
func (instanceApi *InstanceApi) UpdateInstance(c *gin.Context) {
	var instance op.Instance
	_ = c.ShouldBindJSON(&instance)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Platform":{utils.NotEmpty()},
          "Env":{utils.NotEmpty()},
          "Host":{utils.NotEmpty()},
          "User":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(instance, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := instanceService.UpdateInstance(instance); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInstance 用id查询Instance
// @Tags Instance
// @Summary 用id查询Instance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query op.Instance true "用id查询Instance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /instance/findInstance [get]
func (instanceApi *InstanceApi) FindInstance(c *gin.Context) {
	var instance op.Instance
	_ = c.ShouldBindQuery(&instance)
	if reinstance, err := instanceService.GetInstance(instance.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinstance": reinstance}, c)
	}
}

// GetInstanceList 分页获取Instance列表
// @Tags Instance
// @Summary 分页获取Instance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query opReq.InstanceSearch true "分页获取Instance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /instance/getInstanceList [get]
func (instanceApi *InstanceApi) GetInstanceList(c *gin.Context) {
	var pageInfo opReq.InstanceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := instanceService.GetInstanceInfoList(pageInfo); err != nil {
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
