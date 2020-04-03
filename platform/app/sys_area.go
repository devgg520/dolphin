// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_area.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysAreaAdd api implementation
// @Summary 添加区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysArea false "区域信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/area/add [post]
func SysAreaAdd(ctx *Context) {
	var payload model.SysArea
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

// SysAreaDel api implementation
// @Summary 删除区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_area body model.SysArea false "区域"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/area/del [delete]
func SysAreaDel(ctx *Context) {
	var payload model.SysArea
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Table(new(model.SysArea)).In("id", payload.ID.String).Update(map[string]interface{}{
		"update_time": null.TimeFromPtr(time.Now().Value()),
		"update_by":   null.StringFrom(ctx.GetToken().GetUserID()),
		"del_flag":    null.IntFrom(1),
	})
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAreaUpdate api implementation
// @Summary 更新区域
// @Tags 区域
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "区域信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/area/update [put]
func SysAreaUpdate(ctx *Context) {
	var payload model.SysRole
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

// SysAreaPage api implementation
// @Summary 区域分页查询
// @Tags 区域
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/area/page [get]
func SysAreaPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_area", "page", "sys_area", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAreaGet api implementation
// @Summary 获取区域信息
// @Tags 区域
// @Param Authorization header string false "认证令牌"
// @Param id query string false "区域id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/area/get [get]
func SysAreaGet(ctx *Context) {
	var entity model.SysArea
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
