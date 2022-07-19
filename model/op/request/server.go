package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/op"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ServerSearch struct{
    op.Server
    request.PageInfo
}
