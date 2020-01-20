// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_attachment.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/time"
)

// SysAttachmentAdd api implementation
// @Summary 添加附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "附件信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/add [post]
func SysAttachmentAdd(ctx *Context) {
	var form model.SysRole
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	form.CreateTime = null.TimeFromPtr(time.Now().Value())
	form.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentUpdate api implementation
// @Summary 更新附件
// @Tags 附件
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysRole false "附件信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/update [post]
func SysAttachmentUpdate(ctx *Context) {
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

// SysAttachmentPage api implementation
// @Summary 附件分页查询
// @Tags 附件
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/page [get]
func SysAttachmentPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctx.PageSearch(ctx.DB, "sys_attachment", "page", "sys_attachment", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysAttachmentGet api implementation
// @Summary 获取附件信息
// @Tags 附件
// @Param Authorization header string false "认证令牌"
// @Param id query string false "附件id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/attachment/get [get]
func SysAttachmentGet(ctx *Context) {
	var entity model.SysAttachment
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
