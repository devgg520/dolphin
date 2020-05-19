// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user.go

package app

import (
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/time"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/srv"
)

// SysUserAdd api implementation
// @Summary 添加用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/add [post]
func SysUserAdd(ctx *Context) {
	var payload model.SysUser
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
	payload.SetPassword("123456")
	ret, err := ctx.PlatformDB.Insert(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserDel api implementation
// @Summary 删除用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/del [delete]
func SysUserDel(ctx *Context) {
	var payload model.SysUser
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	ret, err := ctx.PlatformDB.In("id", payload.ID.String).Update(&model.SysUser{
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

// SysUserUpdate api implementation
// @Summary 更新用户
// @Tags 用户
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/update [put]
func SysUserUpdate(ctx *Context) {
	var payload model.SysUser
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.Password.Valid = false
	payload.Salt.Valid = false
	ret, err := ctx.PlatformDB.ID(payload.ID).Update(&payload)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysUserPage api implementation
// @Summary 用户分页查询
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Param page query int false "页码"
// @Param size query int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/page [get]
func SysUserPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetString("org_id")
	q.SetString("cn_org_id")
	q.SetTags()

	if q.GetString("cn_org_id") != "" {
		ids, err := srv.SysUserGetOrgsFromInheritance(ctx.DB, q.GetString("cn_org_id"))
		if err != nil {
			ctx.Fail(err)
			return
		}
		q.SetString("cn_org_id", template.HTML(strings.Join(ids, ",")))()
	}

	ret, err := ctx.PageSearch(ctx.PlatformDB, "sys_user", "page", "sys_user", q.Value())
	if err != nil {
		ctx.Fail(err)
		return
	}

	// combile from srv sql
	uids := []string{}
	uorgs := []string{}
	for _, v := range ret.Data {
		if v["id"] != nil {
			uids = append(uids, fmt.Sprintf("'%v'", v["id"]))
		}
		if v["org_id"] != nil {
			uorgs = append(uorgs, fmt.Sprintf("'%v'", v["org_id"]))
		}
	}

	roles, err := srv.SysUserGetUserRolesByUID(ctx.DB, strings.Join(uids, ","))
	if err != nil {
		ctx.Fail(err)
		return
	}

	orgs, err := srv.SysUserGetUserOrgsByUID(ctx.DB, strings.Join(uorgs, ","))
	if err != nil {
		ctx.Fail(err)
		return
	}

	for i := range ret.Data {
		for _, v := range roles {
			if v["user_id"].NullString().String == fmt.Sprintf("%v", ret.Data[i]["id"]) {
				ret.Data[i]["role_name"] = v["role_name"].NullString().String
				ret.Data[i]["user_role"] = v["user_role"].NullString().String
			}
		}
		for _, v := range orgs {
			if v["id"].NullString().String == fmt.Sprintf("%v", ret.Data[i]["org_id"]) {
				ret.Data[i]["org_id"] = v["id"].NullString().String
				ret.Data[i]["org_name"] = v["name"].NullString().String
			}
		}
	}

	ret = ret.With(new([]struct {
		ID       null.String `json:"id" xml:"id"`
		Name     null.String `json:"name" xml:"name"`
		NickName null.String `json:"nickname" xml:"nickname"`
		Mobile   null.String `json:"mobile" xml:"mobile"`
		Email    null.String `json:"email" xml:"email"`
		RoleName null.String `json:"role_name" xml:"role_name"`
		UserRole null.String `json:"user_role" xml:"user_role"`
		OrgName  null.String `json:"org_name" xml:"org_name"`
		OrgID    null.String `json:"org_id" xml:"org_id"`
	}))

	if ctx.QueryBool("__export__") {
		cfg := NewBuildExcelConfig(ret.Data)
		cfg.Format = OptionsetsFormat(ctx.DB)
		ctx.SuccessWithExcel(cfg)
		return
	}
	ctx.Success(ret)
}

// SysUserGet api implementation
// @Summary 获取用户信息
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Param id query string false "用户id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/get [get]
func SysUserGet(ctx *Context) {
	var entity model.SysUser
	id := ctx.Query("id")
	_, err := ctx.PlatformDB.ID(id).Get(&entity)
	entity.Password.Valid = false
	entity.Salt.Valid = false
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(entity)
}

// SysUserLogin api implementation
// @Summary 用户认证
// @Tags 用户
// @Accept application/json
// @Param payload body model.Login false "用户信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/login [post]
func SysUserLogin(ctx *Context) {
	var payload, account = model.Login{}, model.SysUser{}
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		ctx.Fail(err)
		return
	}
	if payload.Domain.String == "" {
		reg := regexp.MustCompile("^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$")
		base := reg.FindAllStringSubmatch(ctx.Request.Host, -1)
		payload.Domain = null.StringFrom(base[0][2])
	}
	account.Domain = payload.Domain
	account.Name = payload.Name
	ext, err := ctx.PlatformDB.Where("del_flag = 0 and status = 1").Get(&account)
	if err != nil || !ext || !account.ValidPassword(payload.Password.String) {
		if err == nil {
			err = errors.New("Account doesn't exist or password error")
		}
		logrus.Errorf("SysUserLogin/ValidPassword:%v", err)
		ctx.Fail(err)
		return
	}
	tgr := &oauth2.TokenGenerateRequest{
		UserID:       account.ID.String,
		Domain:       account.Domain.String,
		ClientID:     viper.GetString("oauth.id"),
		ClientSecret: viper.GetString("oauth.secret"),
		Request:      ctx.Request,
	}
	token, err := ctx.OAuth2.Manager.GenerateAccessToken(oauth2.PasswordCredentials, tgr)
	if err != nil {
		logrus.Errorf("SysUserLogin/GenerateAccessToken:%v", err)
		ctx.Fail(err)
		return
	}
	ctx.Success(map[string]interface{}{
		"access_token":  token.GetAccess(),
		"refresh_token": token.GetRefresh(),
	})
}

// SysUserLogout api implementation
// @Summary 用户退出登录
// @Tags 用户
// @Param Authorization header string false "认证令牌"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/user/logout [get]
func SysUserLogout(ctx *Context) {
	err := ctx.OAuth2.Manager.RemoveAccessToken(ctx.GetToken().GetAccess())
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success("successful")
}
