// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_org.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/time"
)

// SysOrgAdd api implementation
// @Summary 添加组织
// @Tags 组织
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "组织信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/org/add [post]
func SysOrgAdd(ctx *Context) {
	var form model.SysRole
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	form.CreateTime = null.TimeFromPtr(time.Now().Value())
	form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	form.UpdateTime = null.TimeFromPtr(time.Now().Value())
	form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOrgUpdate api implementation
// @Summary 更新组织
// @Tags 组织
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "组织信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/org/update [post]
func SysOrgUpdate(ctx *Context) {
	var form model.SysRole
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	form.UpdateTime = null.TimeFromPtr(time.Now().Value())
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOrgPage api implementation
// @Summary 组织分页查询
// @Tags 组织
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/org/page [get]
func SysOrgPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctx.PageSearch(ctx.DB, "sys_org", "page", "sys_org", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOrgGet api implementation
// @Summary 获取组织信息
// @Tags 组织
// @Param Authorization header string false "认证令牌"
// @Param id query string false "组织id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/org/get [get]
func SysOrgGet(ctx *Context) {
	var entity model.SysOrg
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
