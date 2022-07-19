package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
)

type ServerService struct {
}

// CreateServer 创建Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) CreateServer(server op.Server) (err error) {
	err = global.GetGlobalDBByDBName("business").Create(&server).Error
	return err
}

// DeleteServer 删除Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) DeleteServer(server op.Server) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&server).Error
	return err
}

// DeleteServerByIds 批量删除Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) DeleteServerByIds(ids request.IdsReq) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&[]op.Server{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateServer 更新Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) UpdateServer(server op.Server) (err error) {
	err = global.GetGlobalDBByDBName("business").Save(&server).Error
	return err
}

// GetServer 根据id获取Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) GetServer(id uint) (server op.Server, err error) {
	err = global.GetGlobalDBByDBName("business").Where("id = ?", id).First(&server).Error
	return
}

// GetServerInfoList 分页获取Server记录
// Author [piexlmax](https://github.com/piexlmax)
func (serverService *ServerService) GetServerInfoList(info opReq.ServerSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("business").Model(&op.Server{})
	var servers []op.Server
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Env != nil {
		db = db.Where("env = ?", info.Env)
	}
	if info.OutIp != "" {
		db = db.Where("out_ip LIKE ?", "%"+info.OutIp+"%")
	}
	if info.Platform != nil {
		db = db.Where("platform = ?", info.Platform)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&servers).Error
	return servers, total, err
}
