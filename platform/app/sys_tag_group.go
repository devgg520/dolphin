// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tag_group.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysTagGroupAdd api implementation
// @Summary 添加标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body model.SysTagGroup false "标签组信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tag/group/add [post]
func SysTagGroupAdd(ctx *Context) {
	var payload model.SysTagGroup
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFromPtr(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.DelFlag = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupDel api implementation
// @Summary 删除标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body model.SysTagGroup false "标签"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tag/group/del [delete]
func SysTagGroupDel(ctx *Context) {
	var payload model.SysTagGroup
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysTagGroup{
		UpdateTime: null.TimeFromPtr(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupUpdate api implementation
// @Summary 更新标签组
// @Tags 标签组
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag_group body model.SysTagGroup false "标签组信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tag/group/update [put]
func SysTagGroupUpdate(ctx *Context) {
	var payload model.SysTagGroup
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupPage api implementation
// @Summary 标签组分页查询
// @Tags 标签组
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tag/group/page [get]
func SysTagGroupPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_tag_group", "page", "sys_tag_group", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGroupGet api implementation
// @Summary 获取标签组信息
// @Tags 标签组
// @Param Authorization header string false "认证令牌"
// @Param id query string false "标签组id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/tag/group/get [get]
func SysTagGroupGet(ctx *Context) {
	var entity model.SysTagGroup
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
