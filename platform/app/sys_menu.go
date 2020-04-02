// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_menu.go

package app

import (
	"errors"
	"fmt"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysMenuAdd api implementation
// @Summary 添加菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body model.SysMenu false "菜单信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/menu/add [post]
func SysMenuAdd(ctx *Context) {
	var payload model.SysMenu
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFromPtr(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())

	if !payload.Parent.IsZero() {
		parent := model.SysMenu{}
		ext, err := ctx.DB.SqlMapClient("selectone_sys_menu", &map[string]string{"id": payload.Parent.String}).Get(&parent)
		if err != nil || !ext {
			ctx.Fail(err)
			return
		} else if !ext {
			ctx.Fail(errors.New("the record does not exist"))
			return
		}
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%s%s", payload.ID.String, parent.Inheritance.String))
	} else {
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%s|", payload.ID.String))
	}

	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuDel api implementation
// @Summary 删除菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body model.SysMenu false "菜单"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/menu/del [delete]
func SysMenuDel(ctx *Context) {
	var payload model.SysMenu
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.Table(new(model.SysMenu)).In("id", payload.ID.String).Update(map[string]interface{}{
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

// SysMenuUpdate api implementation
// @Summary 更新菜单
// @Tags 菜单
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_menu body model.SysMenu false "菜单信息"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/menu/update [put]
func SysMenuUpdate(ctx *Context) {
	var payload model.SysMenu
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFromPtr(time.Now().Value())

	if !payload.Parent.IsZero() {
		parent := model.SysMenu{}
		ext, err := ctx.DB.SqlMapClient("selectone_sys_menu", &map[string]string{"id": payload.Parent.String}).Get(&parent)
		if err != nil || !ext {
			ctx.Fail(err)
			return
		} else if !ext {
			ctx.Fail(errors.New("the record does not exist"))
			return
		}
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%s%s", payload.ID.String, parent.Inheritance.String))
	} else {
		payload.Inheritance = null.StringFrom(fmt.Sprintf("|%s|", payload.ID.String))
	}

	ret, err := ctx.DB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuList api implementation
// @Summary 菜单分页查询
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/menu/list [get]
func SysMenuList(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_menu", "page", "sys_menu", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuTree api implementation
// @Summary 菜单树形结构
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Response
// @Router /api/sys/menu/tree [get]
func SysMenuTree(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetString("name")
	q.SetTags()
	ret, err := ctx.TreeSearch(ctx.DB, "sys_menu", "tree", "sys_menu", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysMenuGet api implementation
// @Summary 获取菜单信息
// @Tags 菜单
// @Param Authorization header string false "认证令牌"
// @Param id query string false "菜单id"
// @Failure 403 {object} model.Response
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /api/sys/menu/get [get]
func SysMenuGet(ctx *Context) {
	var entity model.SysMenu
	id := ctx.Query("id")
	ret, err := ctx.DB.Id(id).Get(&entity)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
