// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_oauth2.go

package app

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/viper"
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/util"
	"github.com/go-session/cookie"
	"github.com/go-session/session"
	"golang.org/x/oauth2"
)

func init() {
	var hashKey = []byte("FF51A553-72FC-478B-9AEF-93D6F506DE91")
	session.SetStore(
		cookie.NewCookieStore(
			cookie.SetCookieName("cookie_store_id"),
			cookie.SetHashKey(hashKey),
		),
	)
}

// SysOauth2Login api implementation
// @Summary 登录信息
// @Tags OAuth授权
// @Accept application/json
// @Param user body model.SysUser false "用户信息"
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/login [post]
func SysOauth2Login(ctx *Context) {
	store, err := session.Start(nil, ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Request.ParseForm()
	f := ctx.Request.Form
	domain := f.Get("domain")
	username := f.Get("username")
	password := f.Get("password")
	account := model.SysUser{
		Name:   null.StringFrom(username),
		Domain: null.StringFrom(domain),
	}
	ext, err := ctx.engine.PlatformDB.Where("delete_time is null").Get(&account)
	if err != nil {
		ctx.Fail(err)
		return
	}
	if !ext || !account.ValidPassword(password) {
		ctx.Fail(util.ErrInvalidGrant)
		return
	}
	store.Set("LoggedInUserID", account.ID.String)
	store.Set("LoggedInDomain", account.Domain.String)
	store.Save()
	ctx.Redirect(http.StatusFound, viper.GetString("oauth.auth"))
}

// SysOauth2Affirm api implementation
// @Summary 用户授权
// @Tags OAuth授权
// @Accept application/json
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/affirm [post]
func SysOauth2Affirm(ctx *Context) {
	store, err := session.Start(nil, ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Fail(err)
		return
	}
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	ctx.Request.Form = form
	store.Delete("ReturnUri")
	store.Save()
	err = ctx.engine.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Fail(err)
		return
	}
}

// SysOauth2Authorize api implementation
// @Summary 用户授权
// @Tags OAuth授权
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/authorize [get]
func SysOauth2Authorize(ctx *Context) {
	store, err := session.Start(nil, ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Fail(err)
		return
	}
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	ctx.Request.Form = form
	store.Delete("ReturnUri")
	store.Save()
	err = ctx.engine.OAuth2.HandleAuthorizeRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Fail(err)
		return
	}
}

// SysOauth2Token api implementation
// @Summary 获取令牌
// @Tags OAuth授权
// @Accept application/json
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/token [post]
func SysOauth2Token(ctx *Context) {
	err := ctx.engine.OAuth2.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Fail(err)
	}
}

// SysOauth2URL api implementation
// @Summary 授权地址
// @Tags OAuth授权
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/url [get]
func SysOauth2URL(ctx *Context) {
	state := "redirect_uri=" + ctx.Query("redirect_uri") + "&state=" + ctx.Query("state")
	ret := oa2cfg.AuthCodeURL(state)
	ctx.Success(ret)
}

// SysOauth2Oauth2 api implementation
// @Summary 授权回调
// @Tags OAuth授权
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/oauth2 [get]
func SysOauth2Oauth2(ctx *Context) {
	ctx.Request.ParseForm()
	f := ctx.Request.Form
	state := f.Get("state")
	code := f.Get("code")
	if code == "" {
		ctx.Fail(errors.New("Code not found"))
		return
	}
	ret, err := oa2cfg.Exchange(context.Background(), code)
	if err != nil {
		ctx.Fail(err)
		return
	}
	urlState, err := url.Parse(state)
	if err != nil {
		ctx.Fail(err)
		return
	}
	qState := urlState.Query()
	rawRedirect := qState.Get("redirect_uri")
	rawState := qState.Get("state")
	ctx.SetCookie("access_token", ret.AccessToken, 60*60*30, "/", "*", false, true)
	ctx.SetCookie("refresh_token", ret.RefreshToken, 60*60*30, "/", "*", false, true)
	if strings.TrimSpace(rawRedirect) != "" {
		urlRedirect, err := url.Parse(rawRedirect)
		if err != nil {
			ctx.Fail(err)
			return
		}
		redirect := urlRedirect.Path
		ctx.Redirect(http.StatusFound, redirect+"?state="+rawState)
	} else {
		ctx.Redirect(http.StatusFound, "/?state="+rawState)
	}
	ctx.Success(ret)
}

// SysOauth2Refresh api implementation
// @Summary 刷新令牌
// @Tags OAuth授权
// @Failure 403 {object} model.Response
// @Router /api/sys/oauth2/refresh [get]
func SysOauth2Refresh(ctx *Context) {
	refreshtoken, ok := ctx.engine.OAuth2.BearerAuth(ctx.Request)
	if !ok {
		ctx.Fail(util.ErrInvalidAccessToken)
		return
	}
	token := oauth2.Token{}
	token.Expiry = time.Now()
	token.RefreshToken = refreshtoken
	ret, err := oa2cfg.TokenSource(context.Background(), &token).Token()
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}
