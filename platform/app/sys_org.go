// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_org.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysOrgAdd api implementation
// @Summary 添加组织
// @Tags 组织
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysOrg false "组织信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/org/add [post]
func SysOrgAdd(ctx *Context) {
	var payload model.SysOrg
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFrom(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.DelFlag = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOrgDel api implementation
// @Summary 删除组织
// @Tags 组织
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_org body model.SysOrg false "组织"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/org/del [delete]
func SysOrgDel(ctx *Context) {
	var payload model.SysOrg
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysOrg{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
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
// @Router /api/sys/org/update [put]
func SysOrgUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID).Update(&payload)
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
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_org", "page", "sys_org", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysOrgTree api implementation
// @Summary 菜单树形结构
// @Tags 组织
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Response
// @Router /api/sys/org/tree [get]
func SysOrgTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetTags()
	ret, err := ctx.TreeSearch(ctx.DB, "sys_org", "tree", "sys_org", q.Value())
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
	_, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
