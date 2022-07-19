package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
)

type InstanceService struct {
}

// CreateInstance 创建Instance记录
// Author [piexlmax](https://github.com/piexlmax)
func (instanceService *InstanceService) CreateInstance(instance op.Instance) (err error) {
	err = global.GetGlobalDBByDBName("business").Create(&instance).Error
	return err
}

// DeleteInstance 删除Instance记录
// Author [piexlmax](https://github.com/piexlmax)
func (instanceService *InstanceService) DeleteInstance(instance op.Instance) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&instance).Error
	return err
}

// DeleteInstanceByIds 批量删除Instance记录
// Author [piexlmax](https://github.com/piexlmax)
func (instanceService *InstanceService) DeleteInstanceByIds(ids request.IdsReq) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&[]op.Instance{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInstance 更新Instance记录
// Author [piexlmax](https://github.com/piexlmax)
func (instanceService *InstanceService) UpdateInstance(instance op.Instance) (err error) {
	err = global.GetGlobalDBByDBName("business").Save(&instance).Error
	return err
}

// GetInstance 根据id获取Instance记录
// Author [piexlmax](https://github.com/piexlmax)
func (instanceService *InstanceService) GetInstance(id uint) (instance op.Instance, err error) {
	err = global.GetGlobalDBByDBName("business").Where("id = ?", id).First(&instance).Error
	return
}

// GetInstanceInfoList 分页获取Instance记录
// Author [piexlmax](https://github.com/piexlmax)
func (instanceService *InstanceService) GetInstanceInfoList(info opReq.InstanceSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("business").Model(&op.Instance{})
	var instances []op.Instance
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Platform != nil {
		db = db.Where("platform = ?", info.Platform)
	}
	if info.Env != nil {
		db = db.Where("env = ?", info.Env)
	}
	if info.Host != "" {
		db = db.Where("host LIKE ?", "%"+info.Host+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&instances).Error
	return instances, total, err
}
