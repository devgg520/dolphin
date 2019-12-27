// Code generated by dol build. Only Generate by tools if not existed.
// source: role.go

package app

import (
	"github.com/2637309949/dolphin/cli/platform/model"

	"github.com/2637309949/dolphin/cli/packages/gin/binding"
	"github.com/2637309949/dolphin/cli/packages/null"
)

// RoleAdd api implementation
// @Summary 添加角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param role body model.Role false "角色信息"
// @Failure 403 {object} model.Response
// @Router /api/role/add [post]
func RoleAdd(ctx *Context) {
	var form model.Role
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	form.ID = null.StringFromUUID()
	ret, err := ctx.DB.Insert(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// RoleUpdate api implementation
// @Summary 更新角色
// @Tags 角色
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param role body model.Role false "角色信息"
// @Failure 403 {object} model.Response
// @Router /api/role/update [post]
func RoleUpdate(ctx *Context) {
	var form model.Role
	if err := ctx.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.ID(form.ID).Update(&form)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// RolePage api implementation
// @Summary 角色分页查询
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Router /api/role/page [get]
func RolePage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page")
	q.SetInt("size")
	ret, err := ctx.PageSearch(ctx.DB, "role", "page", "role", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// RoleGet api implementation
// @Summary 获取角色信息
// @Tags 角色
// @Param Authorization header string false "认证令牌"
// @Param id query string false "角色id"
// @Failure 403 {object} model.Response
// @Router /api/role/get [get]
func RoleGet(ctx *Context) {
	var entity model.Role
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
