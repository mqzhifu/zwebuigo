package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectSearch struct{
    op.Project
    request.PageInfo
}
