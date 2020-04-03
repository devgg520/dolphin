// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_dingtalk.go

package app

import (
	"github.com/2637309949/dolphin/platform/srv"
)

// SysDingtalkOauth2 api implementation
// @Summary 授权回调
// @Tags 钉钉
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/dingtalk/oauth2 [get]
func SysDingtalkOauth2(ctx *Context) {
	q := ctx.TypeQuery()
	ret, err := srv.SysDingtalkAction(q)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
