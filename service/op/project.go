package op

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	opReq "github.com/flipped-aurora/gin-vue-admin/server/model/op/request"
)

type ProjectService struct {
}

// CreateProject 创建Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) CreateProject(project op.Project) (err error) {
	err = global.GetGlobalDBByDBName("business").Create(&project).Error
	return err
}

// DeleteProject 删除Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) DeleteProject(project op.Project) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&project).Error
	return err
}

// DeleteProjectByIds 批量删除Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.GetGlobalDBByDBName("business").Delete(&[]op.Project{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProject 更新Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) UpdateProject(project op.Project) (err error) {
	err = global.GetGlobalDBByDBName("business").Save(&project).Error
	return err
}

// GetProject 根据id获取Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) GetProject(id uint) (project op.Project, err error) {
	err = global.GetGlobalDBByDBName("business").Where("id = ?", id).First(&project).Error
	return
}

// GetProjectInfoList 分页获取Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) GetProjectInfoList(info opReq.ProjectSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("business").Model(&op.Project{})
	var projects []op.Project
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projects).Error
	return projects, total, err
}
