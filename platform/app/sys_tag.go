// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_tag.go

package app

import (
	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/platform/model"
)

// SysTagAdd api implementation
// @Summary 添加标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body model.SysTag false "标签信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tag/add [post]
func SysTagAdd(ctx *Context) {
	var payload model.SysTag
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
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
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagDel api implementation
// @Summary 删除标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body model.SysUserTemplate false "标签"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tag/del [delete]
func SysTagDel(ctx *Context) {
	var payload model.SysUserTemplate
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysUserTemplate{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagUpdate api implementation
// @Summary 更新标签
// @Tags 标签
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_tag body model.SysTag false "标签信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tag/update [put]
func SysTagUpdate(ctx *Context) {
	var payload model.SysTag
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagPage api implementation
// @Summary 标签分页查询
// @Tags 标签
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tag/page [get]
func SysTagPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("sys_tag_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_tag", "page", "sys_tag", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysTagGet api implementation
// @Summary 获取标签信息
// @Tags 标签
// @Param Authorization header string false "认证令牌"
// @Param id query string false "标签id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/tag/get [get]
func SysTagGet(ctx *Context) {
	var entity model.SysTag
	id := ctx.Query("id")
	_, err := ctx.DB.ID(id).Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}
