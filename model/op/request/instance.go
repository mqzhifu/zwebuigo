package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type InstanceSearch struct{
    op.Instance
    request.PageInfo
}
