// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/rpc"
	"github.com/2637309949/dolphin/platform/rpc/proto"

	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/util"
)

// Name project
var Name = "platform"

// SysAppFun defined
type SysAppFun struct {
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Tree,
	Get func(ctx *Context)
}

// NewSysAppFun defined
func NewSysAppFun() *SysAppFun {
	ctr := &SysAppFun{}
	ctr.Add = SysAppFunAdd
	ctr.BatchAdd = SysAppFunBatchAdd
	ctr.Del = SysAppFunDel
	ctr.BatchDel = SysAppFunBatchDel
	ctr.Update = SysAppFunUpdate
	ctr.BatchUpdate = SysAppFunBatchUpdate
	ctr.Page = SysAppFunPage
	ctr.Tree = SysAppFunTree
	ctr.Get = SysAppFunGet
	return ctr
}

// SysAppFunRoutes defined
func SysAppFunRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/app/fun/add", Auth, SysAppFunInstance.Add)
	group.Handle("POST", "/sys/app/fun/batch_add", Auth, SysAppFunInstance.BatchAdd)
	group.Handle("DELETE", "/sys/app/fun/del", Auth, SysAppFunInstance.Del)
	group.Handle("DELETE", "/sys/app/fun/batch_del", Auth, SysAppFunInstance.BatchDel)
	group.Handle("PUT", "/sys/app/fun/update", Auth, SysAppFunInstance.Update)
	group.Handle("PUT", "/sys/app/fun/batch_update", Auth, SysAppFunInstance.BatchUpdate)
	group.Handle("GET", "/sys/app/fun/page", Auth, SysAppFunInstance.Page)
	group.Handle("GET", "/sys/app/fun/tree", Auth, SysAppFunInstance.Tree)
	group.Handle("GET", "/sys/app/fun/get", Auth, SysAppFunInstance.Get)
}

// SysAppFunInstance defined
var SysAppFunInstance = NewSysAppFun()

// SysArea defined
type SysArea struct {
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get func(ctx *Context)
}

// NewSysArea defined
func NewSysArea() *SysArea {
	ctr := &SysArea{}
	ctr.Add = SysAreaAdd
	ctr.BatchAdd = SysAreaBatchAdd
	ctr.Del = SysAreaDel
	ctr.BatchDel = SysAreaBatchDel
	ctr.Update = SysAreaUpdate
	ctr.BatchUpdate = SysAreaBatchUpdate
	ctr.Page = SysAreaPage
	ctr.Get = SysAreaGet
	return ctr
}

// SysAreaRoutes defined
func SysAreaRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/area/add", Auth, SysAreaInstance.Add)
	group.Handle("POST", "/sys/area/batch_add", Auth, SysAreaInstance.BatchAdd)
	group.Handle("DELETE", "/sys/area/del", Auth, SysAreaInstance.Del)
	group.Handle("DELETE", "/sys/area/batch_del", Auth, SysAreaInstance.BatchDel)
	group.Handle("PUT", "/sys/area/update", Auth, SysAreaInstance.Update)
	group.Handle("PUT", "/sys/area/batch_update", Auth, SysAreaInstance.BatchUpdate)
	group.Handle("GET", "/sys/area/page", Auth, SysAreaInstance.Page)
	group.Handle("GET", "/sys/area/get", Auth, SysAreaInstance.Get)
}

// SysAreaInstance defined
var SysAreaInstance = NewSysArea()

// SysAttachment defined
type SysAttachment struct {
	Add,
	BatchAdd,
	Upload,
	Export,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get func(ctx *Context)
}

// NewSysAttachment defined
func NewSysAttachment() *SysAttachment {
	ctr := &SysAttachment{}
	ctr.Add = SysAttachmentAdd
	ctr.BatchAdd = SysAttachmentBatchAdd
	ctr.Upload = SysAttachmentUpload
	ctr.Export = SysAttachmentExport
	ctr.Del = SysAttachmentDel
	ctr.BatchDel = SysAttachmentBatchDel
	ctr.Update = SysAttachmentUpdate
	ctr.BatchUpdate = SysAttachmentBatchUpdate
	ctr.Page = SysAttachmentPage
	ctr.Get = SysAttachmentGet
	return ctr
}

// SysAttachmentRoutes defined
func SysAttachmentRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/attachment/add", Auth, SysAttachmentInstance.Add)
	group.Handle("POST", "/sys/attachment/batch_add", Auth, SysAttachmentInstance.BatchAdd)
	group.Handle("POST", "/sys/attachment/upload", Auth, SysAttachmentInstance.Upload)
	group.Handle("GET", "/sys/attachment/export", SysAttachmentInstance.Export)
	group.Handle("DELETE", "/sys/attachment/del", Auth, SysAttachmentInstance.Del)
	group.Handle("POST", "/sys/attachment/batch_del", Auth, SysAttachmentInstance.BatchDel)
	group.Handle("PUT", "/sys/attachment/update", Auth, SysAttachmentInstance.Update)
	group.Handle("POST", "/sys/attachment/batch_update", Auth, SysAttachmentInstance.BatchUpdate)
	group.Handle("GET", "/sys/attachment/page", Auth, SysAttachmentInstance.Page)
	group.Handle("GET", "/sys/attachment/get", Auth, SysAttachmentInstance.Get)
}

// SysAttachmentInstance defined
var SysAttachmentInstance = NewSysAttachment()

// SysCas defined
type SysCas struct {
	Login,
	Logout,
	Affirm,
	Authorize,
	Token,
	URL,
	Oauth2,
	Refresh,
	Check,
	Profile,
	Qrcode func(ctx *Context)
}

// NewSysCas defined
func NewSysCas() *SysCas {
	ctr := &SysCas{}
	ctr.Login = SysCasLogin
	ctr.Logout = SysCasLogout
	ctr.Affirm = SysCasAffirm
	ctr.Authorize = SysCasAuthorize
	ctr.Token = SysCasToken
	ctr.URL = SysCasURL
	ctr.Oauth2 = SysCasOauth2
	ctr.Refresh = SysCasRefresh
	ctr.Check = SysCasCheck
	ctr.Profile = SysCasProfile
	ctr.Qrcode = SysCasQrcode
	return ctr
}

// SysCasRoutes defined
func SysCasRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/cas/login", SysCasInstance.Login)
	group.Handle("GET", "/sys/cas/logout", SysCasInstance.Logout)
	group.Handle("POST", "/sys/cas/affirm", SysCasInstance.Affirm)
	group.Handle("GET", "/sys/cas/authorize", SysCasInstance.Authorize)
	group.Handle("POST", "/sys/cas/token", SysCasInstance.Token)
	group.Handle("GET", "/sys/cas/url", SysCasInstance.URL)
	group.Handle("GET", "/sys/cas/oauth2", SysCasInstance.Oauth2)
	group.Handle("GET", "/sys/cas/refresh", SysCasInstance.Refresh)
	group.Handle("GET", "/sys/cas/check", SysCasInstance.Check)
	group.Handle("GET", "/sys/cas/profile", Auth, SysCasInstance.Profile)
	group.Handle("GET", "/sys/cas/qrcode", SysCasInstance.Qrcode)
}

// SysCasInstance defined
var SysCasInstance = NewSysCas()

// SysClient defined
type SysClient struct {
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get func(ctx *Context)
}

// NewSysClient defined
func NewSysClient() *SysClient {
	ctr := &SysClient{}
	ctr.Add = SysClientAdd
	ctr.BatchAdd = SysClientBatchAdd
	ctr.Del = SysClientDel
	ctr.BatchDel = SysClientBatchDel
	ctr.Update = SysClientUpdate
	ctr.BatchUpdate = SysClientBatchUpdate
	ctr.Page = SysClientPage
	ctr.Get = SysClientGet
	return ctr
}

// SysClientRoutes defined
func SysClientRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/client/add", Auth, SysClientInstance.Add)
	group.Handle("POST", "/sys/client/batch_add", Auth, SysClientInstance.BatchAdd)
	group.Handle("DELETE", "/sys/client/del", Auth, SysClientInstance.Del)
	group.Handle("DELETE", "/sys/client/batch_del", Auth, SysClientInstance.BatchDel)
	group.Handle("PUT", "/sys/client/update", Auth, SysClientInstance.Update)
	group.Handle("PUT", "/sys/client/batch_update", Auth, SysClientInstance.BatchUpdate)
	group.Handle("GET", "/sys/client/page", Auth, SysClientInstance.Page)
	group.Handle("GET", "/sys/client/get", Auth, SysClientInstance.Get)
}

// SysClientInstance defined
var SysClientInstance = NewSysClient()

// SysComment defined
type SysComment struct {
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get func(ctx *Context)
}

// NewSysComment defined
func NewSysComment() *SysComment {
	ctr := &SysComment{}
	ctr.Add = SysCommentAdd
	ctr.BatchAdd = SysCommentBatchAdd
	ctr.Del = SysCommentDel
	ctr.BatchDel = SysCommentBatchDel
	ctr.Update = SysCommentUpdate
	ctr.BatchUpdate = SysCommentBatchUpdate
	ctr.Page = SysCommentPage
	ctr.Get = SysCommentGet
	return ctr
}

// SysCommentRoutes defined
func SysCommentRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/comment/add", Auth, SysCommentInstance.Add)
	group.Handle("POST", "/sys/comment/batch_add", Auth, SysCommentInstance.BatchAdd)
	group.Handle("DELETE", "/sys/comment/del", Auth, SysCommentInstance.Del)
	group.Handle("DELETE", "/sys/comment/batch_del", Auth, SysCommentInstance.BatchDel)
	group.Handle("PUT", "/sys/comment/update", Auth, SysCommentInstance.Update)
	group.Handle("PUT", "/sys/comment/batch_update", Auth, SysCommentInstance.BatchUpdate)
	group.Handle("GET", "/sys/comment/page", Auth, SysCommentInstance.Page)
	group.Handle("GET", "/sys/comment/get", Auth, SysCommentInstance.Get)
}

// SysCommentInstance defined
var SysCommentInstance = NewSysComment()

// SysDataPermission defined
type SysDataPermission struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysDataPermission defined
func NewSysDataPermission() *SysDataPermission {
	ctr := &SysDataPermission{}
	ctr.Add = SysDataPermissionAdd
	ctr.Del = SysDataPermissionDel
	ctr.Update = SysDataPermissionUpdate
	ctr.Page = SysDataPermissionPage
	ctr.Get = SysDataPermissionGet
	return ctr
}

// SysDataPermissionRoutes defined
func SysDataPermissionRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/data/permission/add", Auth, SysDataPermissionInstance.Add)
	group.Handle("DELETE", "/sys/data/permission/del", Auth, SysDataPermissionInstance.Del)
	group.Handle("PUT", "/sys/data/permission/update", Auth, SysDataPermissionInstance.Update)
	group.Handle("GET", "/sys/data/permission/page", Auth, SysDataPermissionInstance.Page)
	group.Handle("GET", "/sys/data/permission/get", Auth, SysDataPermissionInstance.Get)
}

// SysDataPermissionInstance defined
var SysDataPermissionInstance = NewSysDataPermission()

// Debug defined
type Debug struct {
	Pprof,
	Heap,
	Goroutine,
	Allocs,
	Block,
	Threadcreate,
	Cmdline,
	Profile,
	Symbol,
	Trace,
	Mutex func(ctx *Context)
}

// NewDebug defined
func NewDebug() *Debug {
	ctr := &Debug{}
	ctr.Pprof = DebugPprof
	ctr.Heap = DebugHeap
	ctr.Goroutine = DebugGoroutine
	ctr.Allocs = DebugAllocs
	ctr.Block = DebugBlock
	ctr.Threadcreate = DebugThreadcreate
	ctr.Cmdline = DebugCmdline
	ctr.Profile = DebugProfile
	ctr.Symbol = DebugSymbol
	ctr.Trace = DebugTrace
	ctr.Mutex = DebugMutex
	return ctr
}

// DebugRoutes defined
func DebugRoutes(engine *Engine) {
	group := engine.Group("/debug")
	group.Handle("GET", "/pprof/", Auth, Roles("X8e6D3y60K"), DebugInstance.Pprof)
	group.Handle("GET", "/pprof/heap", Auth, Roles("X8e6D3y60K"), DebugInstance.Heap)
	group.Handle("GET", "/pprof/goroutine", Auth, Roles("X8e6D3y60K"), DebugInstance.Goroutine)
	group.Handle("GET", "/pprof/allocs", Auth, Roles("X8e6D3y60K"), DebugInstance.Allocs)
	group.Handle("GET", "/pprof/block", Auth, Roles("X8e6D3y60K"), DebugInstance.Block)
	group.Handle("GET", "/pprof/threadcreate", Auth, Roles("X8e6D3y60K"), DebugInstance.Threadcreate)
	group.Handle("GET", "/pprof/cmdline", Auth, Roles("X8e6D3y60K"), DebugInstance.Cmdline)
	group.Handle("GET", "/pprof/profile", Auth, Roles("X8e6D3y60K"), DebugInstance.Profile)
	group.Handle("GET,POST", "/pprof/symbol", Auth, Roles("X8e6D3y60K"), DebugInstance.Symbol)
	group.Handle("GET", "/pprof/trace", Auth, Roles("X8e6D3y60K"), DebugInstance.Trace)
	group.Handle("GET", "/pprof/mutex", Auth, Roles("X8e6D3y60K"), DebugInstance.Mutex)
}

// DebugInstance defined
var DebugInstance = NewDebug()

// SysDingtalk defined
type SysDingtalk struct {
	Oauth2 func(ctx *Context)
}

// NewSysDingtalk defined
func NewSysDingtalk() *SysDingtalk {
	ctr := &SysDingtalk{}
	ctr.Oauth2 = SysDingtalkOauth2
	return ctr
}

// SysDingtalkRoutes defined
func SysDingtalkRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/dingtalk/oauth2", SysDingtalkInstance.Oauth2)
}

// SysDingtalkInstance defined
var SysDingtalkInstance = NewSysDingtalk()

// SysDomain defined
type SysDomain struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysDomain defined
func NewSysDomain() *SysDomain {
	ctr := &SysDomain{}
	ctr.Add = SysDomainAdd
	ctr.Del = SysDomainDel
	ctr.Update = SysDomainUpdate
	ctr.Page = SysDomainPage
	ctr.Get = SysDomainGet
	return ctr
}

// SysDomainRoutes defined
func SysDomainRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/domain/add", Auth, SysDomainInstance.Add)
	group.Handle("DELETE", "/sys/domain/del", Auth, SysDomainInstance.Del)
	group.Handle("PUT", "/sys/domain/update", Auth, SysDomainInstance.Update)
	group.Handle("GET", "/sys/domain/page", Auth, SysDomainInstance.Page)
	group.Handle("GET", "/sys/domain/get", Auth, SysDomainInstance.Get)
}

// SysDomainInstance defined
var SysDomainInstance = NewSysDomain()

// SysMenu defined
type SysMenu struct {
	Add,
	Del,
	BatchDel,
	Update,
	Sidebar,
	Page,
	Tree,
	Get func(ctx *Context)
}

// NewSysMenu defined
func NewSysMenu() *SysMenu {
	ctr := &SysMenu{}
	ctr.Add = SysMenuAdd
	ctr.Del = SysMenuDel
	ctr.BatchDel = SysMenuBatchDel
	ctr.Update = SysMenuUpdate
	ctr.Sidebar = SysMenuSidebar
	ctr.Page = SysMenuPage
	ctr.Tree = SysMenuTree
	ctr.Get = SysMenuGet
	return ctr
}

// SysMenuRoutes defined
func SysMenuRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/menu/add", Auth, SysMenuInstance.Add)
	group.Handle("DELETE", "/sys/menu/del", Auth, SysMenuInstance.Del)
	group.Handle("DELETE", "/sys/menu/batch_del", Auth, SysMenuInstance.BatchDel)
	group.Handle("PUT", "/sys/menu/update", Auth, SysMenuInstance.Update)
	group.Handle("GET", "/sys/menu/sidebar", Auth, SysMenuInstance.Sidebar)
	group.Handle("GET", "/sys/menu/page", Auth, SysMenuInstance.Page)
	group.Handle("GET", "/sys/menu/tree", Auth, SysMenuInstance.Tree)
	group.Handle("GET", "/sys/menu/get", Auth, SysMenuInstance.Get)
}

// SysMenuInstance defined
var SysMenuInstance = NewSysMenu()

// SysNotification defined
type SysNotification struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysNotification defined
func NewSysNotification() *SysNotification {
	ctr := &SysNotification{}
	ctr.Add = SysNotificationAdd
	ctr.Del = SysNotificationDel
	ctr.Update = SysNotificationUpdate
	ctr.Page = SysNotificationPage
	ctr.Get = SysNotificationGet
	return ctr
}

// SysNotificationRoutes defined
func SysNotificationRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/notification/add", Auth, SysNotificationInstance.Add)
	group.Handle("DELETE", "/sys/notification/del", Auth, SysNotificationInstance.Del)
	group.Handle("PUT", "/sys/notification/update", Auth, SysNotificationInstance.Update)
	group.Handle("GET", "/sys/notification/page", Auth, SysNotificationInstance.Page)
	group.Handle("GET", "/sys/notification/get", Auth, SysNotificationInstance.Get)
}

// SysNotificationInstance defined
var SysNotificationInstance = NewSysNotification()

// SysOptionset defined
type SysOptionset struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysOptionset defined
func NewSysOptionset() *SysOptionset {
	ctr := &SysOptionset{}
	ctr.Add = SysOptionsetAdd
	ctr.Del = SysOptionsetDel
	ctr.BatchDel = SysOptionsetBatchDel
	ctr.Update = SysOptionsetUpdate
	ctr.Page = SysOptionsetPage
	ctr.Get = SysOptionsetGet
	return ctr
}

// SysOptionsetRoutes defined
func SysOptionsetRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/optionset/add", Auth, SysOptionsetInstance.Add)
	group.Handle("DELETE", "/sys/optionset/del", Auth, SysOptionsetInstance.Del)
	group.Handle("DELETE", "/sys/optionset/batch_del", Auth, SysOptionsetInstance.BatchDel)
	group.Handle("PUT", "/sys/optionset/update", Auth, SysOptionsetInstance.Update)
	group.Handle("GET", "/sys/optionset/page", Auth, SysOptionsetInstance.Page)
	group.Handle("GET", "/sys/optionset/get", Auth, SysOptionsetInstance.Get)
}

// SysOptionsetInstance defined
var SysOptionsetInstance = NewSysOptionset()

// SysOrg defined
type SysOrg struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Tree,
	Get func(ctx *Context)
}

// NewSysOrg defined
func NewSysOrg() *SysOrg {
	ctr := &SysOrg{}
	ctr.Add = SysOrgAdd
	ctr.Del = SysOrgDel
	ctr.BatchDel = SysOrgBatchDel
	ctr.Update = SysOrgUpdate
	ctr.Page = SysOrgPage
	ctr.Tree = SysOrgTree
	ctr.Get = SysOrgGet
	return ctr
}

// SysOrgRoutes defined
func SysOrgRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/org/add", Auth, SysOrgInstance.Add)
	group.Handle("DELETE", "/sys/org/del", Auth, SysOrgInstance.Del)
	group.Handle("DELETE", "/sys/org/batch_del", Auth, SysOrgInstance.BatchDel)
	group.Handle("PUT", "/sys/org/update", Auth, SysOrgInstance.Update)
	group.Handle("GET", "/sys/org/page", Auth, SysOrgInstance.Page)
	group.Handle("GET", "/sys/org/tree", Auth, SysOrgInstance.Tree)
	group.Handle("GET", "/sys/org/get", Auth, SysOrgInstance.Get)
}

// SysOrgInstance defined
var SysOrgInstance = NewSysOrg()

// SysPermission defined
type SysPermission struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysPermission defined
func NewSysPermission() *SysPermission {
	ctr := &SysPermission{}
	ctr.Add = SysPermissionAdd
	ctr.Del = SysPermissionDel
	ctr.Update = SysPermissionUpdate
	ctr.Page = SysPermissionPage
	ctr.Get = SysPermissionGet
	return ctr
}

// SysPermissionRoutes defined
func SysPermissionRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/permission/add", Auth, SysPermissionInstance.Add)
	group.Handle("DELETE", "/sys/permission/del", Auth, SysPermissionInstance.Del)
	group.Handle("PUT", "/sys/permission/update", Auth, SysPermissionInstance.Update)
	group.Handle("GET", "/sys/permission/page", Auth, SysPermissionInstance.Page)
	group.Handle("GET", "/sys/permission/get", Auth, SysPermissionInstance.Get)
}

// SysPermissionInstance defined
var SysPermissionInstance = NewSysPermission()

// SysRole defined
type SysRole struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	RoleMenuTree,
	RoleAppFunTree,
	Get func(ctx *Context)
}

// NewSysRole defined
func NewSysRole() *SysRole {
	ctr := &SysRole{}
	ctr.Add = SysRoleAdd
	ctr.Del = SysRoleDel
	ctr.BatchDel = SysRoleBatchDel
	ctr.Update = SysRoleUpdate
	ctr.Page = SysRolePage
	ctr.RoleMenuTree = SysRoleRoleMenuTree
	ctr.RoleAppFunTree = SysRoleRoleAppFunTree
	ctr.Get = SysRoleGet
	return ctr
}

// SysRoleRoutes defined
func SysRoleRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/role/add", Auth, SysRoleInstance.Add)
	group.Handle("DELETE", "/sys/role/del", Auth, SysRoleInstance.Del)
	group.Handle("DELETE", "/sys/role/batch_del", Auth, SysRoleInstance.BatchDel)
	group.Handle("PUT", "/sys/role/update", Auth, SysRoleInstance.Update)
	group.Handle("GET", "/sys/role/page", Auth, SysRoleInstance.Page)
	group.Handle("GET", "/sys/role/role_menu_tree", Auth, SysRoleInstance.RoleMenuTree)
	group.Handle("GET", "/sys/role/role_app_fun_tree", Auth, SysRoleInstance.RoleAppFunTree)
	group.Handle("GET", "/sys/role/get", Auth, SysRoleInstance.Get)
}

// SysRoleInstance defined
var SysRoleInstance = NewSysRole()

// SysRoleMenu defined
type SysRoleMenu struct {
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysRoleMenu defined
func NewSysRoleMenu() *SysRoleMenu {
	ctr := &SysRoleMenu{}
	ctr.Add = SysRoleMenuAdd
	ctr.BatchAdd = SysRoleMenuBatchAdd
	ctr.Del = SysRoleMenuDel
	ctr.BatchDel = SysRoleMenuBatchDel
	ctr.Update = SysRoleMenuUpdate
	ctr.Page = SysRoleMenuPage
	ctr.Get = SysRoleMenuGet
	return ctr
}

// SysRoleMenuRoutes defined
func SysRoleMenuRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/role/menu/add", Auth, SysRoleMenuInstance.Add)
	group.Handle("POST", "/sys/role/menu/batch_add", Auth, SysRoleMenuInstance.BatchAdd)
	group.Handle("DELETE", "/sys/role/menu/del", Auth, SysRoleMenuInstance.Del)
	group.Handle("DELETE", "/sys/role/menu/batch_del", Auth, SysRoleMenuInstance.BatchDel)
	group.Handle("PUT", "/sys/role/menu/update", Auth, SysRoleMenuInstance.Update)
	group.Handle("GET", "/sys/role/menu/page", Auth, SysRoleMenuInstance.Page)
	group.Handle("GET", "/sys/role/menu/get", Auth, SysRoleMenuInstance.Get)
}

// SysRoleMenuInstance defined
var SysRoleMenuInstance = NewSysRoleMenu()

// SysSchedule defined
type SysSchedule struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysSchedule defined
func NewSysSchedule() *SysSchedule {
	ctr := &SysSchedule{}
	ctr.Add = SysScheduleAdd
	ctr.Del = SysScheduleDel
	ctr.BatchDel = SysScheduleBatchDel
	ctr.Update = SysScheduleUpdate
	ctr.Page = SysSchedulePage
	ctr.Get = SysScheduleGet
	return ctr
}

// SysScheduleRoutes defined
func SysScheduleRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/schedule/add", Auth, SysScheduleInstance.Add)
	group.Handle("DELETE", "/sys/schedule/del", Auth, SysScheduleInstance.Del)
	group.Handle("DELETE", "/sys/schedule/batch_del", Auth, SysScheduleInstance.BatchDel)
	group.Handle("PUT", "/sys/schedule/update", Auth, SysScheduleInstance.Update)
	group.Handle("GET", "/sys/schedule/page", Auth, SysScheduleInstance.Page)
	group.Handle("GET", "/sys/schedule/get", Auth, SysScheduleInstance.Get)
}

// SysScheduleInstance defined
var SysScheduleInstance = NewSysSchedule()

// SysScheduleHistory defined
type SysScheduleHistory struct {
	Page func(ctx *Context)
}

// NewSysScheduleHistory defined
func NewSysScheduleHistory() *SysScheduleHistory {
	ctr := &SysScheduleHistory{}
	ctr.Page = SysScheduleHistoryPage
	return ctr
}

// SysScheduleHistoryRoutes defined
func SysScheduleHistoryRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/schedule/history/page", Auth, SysScheduleHistoryInstance.Page)
}

// SysScheduleHistoryInstance defined
var SysScheduleHistoryInstance = NewSysScheduleHistory()

// SysScheduling defined
type SysScheduling struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysScheduling defined
func NewSysScheduling() *SysScheduling {
	ctr := &SysScheduling{}
	ctr.Add = SysSchedulingAdd
	ctr.Del = SysSchedulingDel
	ctr.Update = SysSchedulingUpdate
	ctr.Page = SysSchedulingPage
	ctr.Get = SysSchedulingGet
	return ctr
}

// SysSchedulingRoutes defined
func SysSchedulingRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/scheduling/add", Auth, SysSchedulingInstance.Add)
	group.Handle("DELETE", "/sys/scheduling/del", Auth, SysSchedulingInstance.Del)
	group.Handle("PUT", "/sys/scheduling/update", Auth, SysSchedulingInstance.Update)
	group.Handle("GET", "/sys/scheduling/page", Auth, SysSchedulingInstance.Page)
	group.Handle("GET", "/sys/scheduling/get", Auth, SysSchedulingInstance.Get)
}

// SysSchedulingInstance defined
var SysSchedulingInstance = NewSysScheduling()

// SysSetting defined
type SysSetting struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysSetting defined
func NewSysSetting() *SysSetting {
	ctr := &SysSetting{}
	ctr.Add = SysSettingAdd
	ctr.Del = SysSettingDel
	ctr.BatchDel = SysSettingBatchDel
	ctr.Update = SysSettingUpdate
	ctr.Page = SysSettingPage
	ctr.Get = SysSettingGet
	return ctr
}

// SysSettingRoutes defined
func SysSettingRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/setting/add", Auth, SysSettingInstance.Add)
	group.Handle("DELETE", "/sys/setting/del", Auth, SysSettingInstance.Del)
	group.Handle("DELETE", "/sys/setting/batch_del", Auth, SysSettingInstance.BatchDel)
	group.Handle("PUT", "/sys/setting/update", Auth, SysSettingInstance.Update)
	group.Handle("GET", "/sys/setting/page", Auth, SysSettingInstance.Page)
	group.Handle("GET", "/sys/setting/get", Auth, SysSettingInstance.Get)
}

// SysSettingInstance defined
var SysSettingInstance = NewSysSetting()

// SysTable defined
type SysTable struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysTable defined
func NewSysTable() *SysTable {
	ctr := &SysTable{}
	ctr.Add = SysTableAdd
	ctr.Del = SysTableDel
	ctr.BatchDel = SysTableBatchDel
	ctr.Update = SysTableUpdate
	ctr.Page = SysTablePage
	ctr.Get = SysTableGet
	return ctr
}

// SysTableRoutes defined
func SysTableRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/table/add", Auth, SysTableInstance.Add)
	group.Handle("DELETE", "/sys/table/del", Auth, SysTableInstance.Del)
	group.Handle("DELETE", "/sys/table/batch_del", Auth, SysTableInstance.BatchDel)
	group.Handle("PUT", "/sys/table/update", Auth, SysTableInstance.Update)
	group.Handle("GET", "/sys/table/page", Auth, SysTableInstance.Page)
	group.Handle("GET", "/sys/table/get", Auth, SysTableInstance.Get)
}

// SysTableInstance defined
var SysTableInstance = NewSysTable()

// SysTableColumn defined
type SysTableColumn struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysTableColumn defined
func NewSysTableColumn() *SysTableColumn {
	ctr := &SysTableColumn{}
	ctr.Add = SysTableColumnAdd
	ctr.Del = SysTableColumnDel
	ctr.BatchDel = SysTableColumnBatchDel
	ctr.Update = SysTableColumnUpdate
	ctr.Page = SysTableColumnPage
	ctr.Get = SysTableColumnGet
	return ctr
}

// SysTableColumnRoutes defined
func SysTableColumnRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/table/column/add", Auth, SysTableColumnInstance.Add)
	group.Handle("DELETE", "/sys/table/column/del", Auth, SysTableColumnInstance.Del)
	group.Handle("DELETE", "/sys/table/column/batch_del", Auth, SysTableColumnInstance.BatchDel)
	group.Handle("PUT", "/sys/table/column/update", Auth, SysTableColumnInstance.Update)
	group.Handle("GET", "/sys/table/column/page", Auth, SysTableColumnInstance.Page)
	group.Handle("GET", "/sys/table/column/get", Auth, SysTableColumnInstance.Get)
}

// SysTableColumnInstance defined
var SysTableColumnInstance = NewSysTableColumn()

// SysTag defined
type SysTag struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysTag defined
func NewSysTag() *SysTag {
	ctr := &SysTag{}
	ctr.Add = SysTagAdd
	ctr.Del = SysTagDel
	ctr.BatchDel = SysTagBatchDel
	ctr.Update = SysTagUpdate
	ctr.Page = SysTagPage
	ctr.Get = SysTagGet
	return ctr
}

// SysTagRoutes defined
func SysTagRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/tag/add", Auth, SysTagInstance.Add)
	group.Handle("DELETE", "/sys/tag/del", Auth, SysTagInstance.Del)
	group.Handle("DELETE", "/sys/tag/batch_del", Auth, SysTagInstance.BatchDel)
	group.Handle("PUT", "/sys/tag/update", Auth, SysTagInstance.Update)
	group.Handle("GET", "/sys/tag/page", Auth, SysTagInstance.Page)
	group.Handle("GET", "/sys/tag/get", Auth, SysTagInstance.Get)
}

// SysTagInstance defined
var SysTagInstance = NewSysTag()

// SysTagGroup defined
type SysTagGroup struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysTagGroup defined
func NewSysTagGroup() *SysTagGroup {
	ctr := &SysTagGroup{}
	ctr.Add = SysTagGroupAdd
	ctr.Del = SysTagGroupDel
	ctr.Update = SysTagGroupUpdate
	ctr.Page = SysTagGroupPage
	ctr.Get = SysTagGroupGet
	return ctr
}

// SysTagGroupRoutes defined
func SysTagGroupRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/tag/group/add", Auth, SysTagGroupInstance.Add)
	group.Handle("DELETE", "/sys/tag/group/del", Auth, SysTagGroupInstance.Del)
	group.Handle("PUT", "/sys/tag/group/update", Auth, SysTagGroupInstance.Update)
	group.Handle("GET", "/sys/tag/group/page", Auth, SysTagGroupInstance.Page)
	group.Handle("GET", "/sys/tag/group/get", Auth, SysTagGroupInstance.Get)
}

// SysTagGroupInstance defined
var SysTagGroupInstance = NewSysTagGroup()

// SysTracker defined
type SysTracker struct {
	Page,
	Get func(ctx *Context)
}

// NewSysTracker defined
func NewSysTracker() *SysTracker {
	ctr := &SysTracker{}
	ctr.Page = SysTrackerPage
	ctr.Get = SysTrackerGet
	return ctr
}

// SysTrackerRoutes defined
func SysTrackerRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/tracker/page", Auth, SysTrackerInstance.Page)
	group.Handle("GET", "/sys/tracker/get", Auth, SysTrackerInstance.Get)
}

// SysTrackerInstance defined
var SysTrackerInstance = NewSysTracker()

// SysUser defined
type SysUser struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Get,
	Login,
	Logout func(ctx *Context)
}

// NewSysUser defined
func NewSysUser() *SysUser {
	ctr := &SysUser{}
	ctr.Add = SysUserAdd
	ctr.Del = SysUserDel
	ctr.BatchDel = SysUserBatchDel
	ctr.Update = SysUserUpdate
	ctr.Page = SysUserPage
	ctr.Get = SysUserGet
	ctr.Login = SysUserLogin
	ctr.Logout = SysUserLogout
	return ctr
}

// SysUserRoutes defined
func SysUserRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/user/add", Auth, Roles("X8e6D3y60K"), SysUserInstance.Add)
	group.Handle("DELETE", "/sys/user/del", Auth, Roles("X8e6D3y60K"), SysUserInstance.Del)
	group.Handle("DELETE", "/sys/user/batch_del", Auth, Roles("X8e6D3y60K"), SysUserInstance.BatchDel)
	group.Handle("PUT", "/sys/user/update", Auth, Roles("X8e6D3y60K"), SysUserInstance.Update)
	group.Handle("GET", "/sys/user/page", Auth, SysUserInstance.Page)
	group.Handle("GET", "/sys/user/get", Auth, SysUserInstance.Get)
	group.Handle("POST", "/sys/user/login", SysUserInstance.Login)
	group.Handle("GET", "/sys/user/logout", Auth, SysUserInstance.Logout)
}

// SysUserInstance defined
var SysUserInstance = NewSysUser()

// SysUserTemplate defined
type SysUserTemplate struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysUserTemplate defined
func NewSysUserTemplate() *SysUserTemplate {
	ctr := &SysUserTemplate{}
	ctr.Add = SysUserTemplateAdd
	ctr.Del = SysUserTemplateDel
	ctr.Update = SysUserTemplateUpdate
	ctr.Page = SysUserTemplatePage
	ctr.Get = SysUserTemplateGet
	return ctr
}

// SysUserTemplateRoutes defined
func SysUserTemplateRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/user/template/add", Auth, SysUserTemplateInstance.Add)
	group.Handle("DELETE", "/sys/user/template/del", Auth, SysUserTemplateInstance.Del)
	group.Handle("PUT", "/sys/user/template/update", Auth, SysUserTemplateInstance.Update)
	group.Handle("GET", "/sys/user/template/page", Auth, SysUserTemplateInstance.Page)
	group.Handle("GET", "/sys/user/template/get", Auth, SysUserTemplateInstance.Get)
}

// SysUserTemplateInstance defined
var SysUserTemplateInstance = NewSysUserTemplate()

// SysUserTemplateDetail defined
type SysUserTemplateDetail struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysUserTemplateDetail defined
func NewSysUserTemplateDetail() *SysUserTemplateDetail {
	ctr := &SysUserTemplateDetail{}
	ctr.Add = SysUserTemplateDetailAdd
	ctr.Del = SysUserTemplateDetailDel
	ctr.Update = SysUserTemplateDetailUpdate
	ctr.Page = SysUserTemplateDetailPage
	ctr.Get = SysUserTemplateDetailGet
	return ctr
}

// SysUserTemplateDetailRoutes defined
func SysUserTemplateDetailRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/user/template/detail/add", Auth, SysUserTemplateDetailInstance.Add)
	group.Handle("DELETE", "/sys/user/template/detail/del", Auth, SysUserTemplateDetailInstance.Del)
	group.Handle("PUT", "/sys/user/template/detail/update", Auth, SysUserTemplateDetailInstance.Update)
	group.Handle("GET", "/sys/user/template/detail/page", Auth, SysUserTemplateDetailInstance.Page)
	group.Handle("GET", "/sys/user/template/detail/get", Auth, SysUserTemplateDetailInstance.Get)
}

// SysUserTemplateDetailInstance defined
var SysUserTemplateDetailInstance = NewSysUserTemplateDetail()

// SysWechat defined
type SysWechat struct {
	Oauth2 func(ctx *Context)
}

// NewSysWechat defined
func NewSysWechat() *SysWechat {
	ctr := &SysWechat{}
	ctr.Oauth2 = SysWechatOauth2
	return ctr
}

// SysWechatRoutes defined
func SysWechatRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/wechat/oauth2", SysWechatInstance.Oauth2)
}

// SysWechatInstance defined
var SysWechatInstance = NewSysWechat()

// SysWorker defined
type SysWorker struct {
	Add,
	Get func(ctx *Context)
}

// NewSysWorker defined
func NewSysWorker() *SysWorker {
	ctr := &SysWorker{}
	ctr.Add = SysWorkerAdd
	ctr.Get = SysWorkerGet
	return ctr
}

// SysWorkerRoutes defined
func SysWorkerRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/worker/add", Auth, SysWorkerInstance.Add)
	group.Handle("GET", "/sys/worker/get", Auth, SysWorkerInstance.Get)
}

// SysWorkerInstance defined
var SysWorkerInstance = NewSysWorker()

// ClientSrv defined
func ClientSrvService(engine *Engine) {
	proto.RegisterClientSrvServer(engine.GRPC, &rpc.ClientSrv{})
}

// DomainSrv defined
func DomainSrvService(engine *Engine) {
	proto.RegisterDomainSrvServer(engine.GRPC, &rpc.DomainSrv{})
}

// UserSrv defined
func UserSrvService(engine *Engine) {
	proto.RegisterUserSrvServer(engine.GRPC, &rpc.UserSrv{})
}

// SyncModel defined
func SyncModel() error {
	mseti := App.Manager.MSet()
	mseti.Add(new(model.SysAppFun))
	mseti.Add(new(model.SysArea))
	mseti.Add(new(model.SysAreaTemplate))
	mseti.Add(new(model.SysAreaTemplateDetail))
	mseti.Add(new(model.SysAttachment))
	mseti.Add(new(model.SysClient), "platform")
	mseti.Add(new(model.SysComment))
	mseti.Add(new(model.SysCommentReply))
	mseti.Add(new(model.SysDataPermission))
	mseti.Add(new(model.SysDataPermissionDetail))
	mseti.Add(new(model.SysDomain), "platform")
	mseti.Add(new(model.SysEmailToken))
	mseti.Add(new(model.SysMenu))
	mseti.Add(new(model.SysNotification))
	mseti.Add(new(model.SysOptionset))
	mseti.Add(new(model.SysOrg))
	mseti.Add(new(model.SysPermission))
	mseti.Add(new(model.SysRole))
	mseti.Add(new(model.SysRoleAppFun))
	mseti.Add(new(model.SysRoleDataPermission))
	mseti.Add(new(model.SysRoleMenu))
	mseti.Add(new(model.SysRolePermission))
	mseti.Add(new(model.SysRoleUser))
	mseti.Add(new(model.SysSchedule))
	mseti.Add(new(model.SysScheduleHistory))
	mseti.Add(new(model.SysSetting))
	mseti.Add(new(model.SysTable))
	mseti.Add(new(model.SysTableColUser))
	mseti.Add(new(model.SysTableColumn))
	mseti.Add(new(model.SysTag))
	mseti.Add(new(model.SysTagGroup))
	mseti.Add(new(model.SysTracker))
	mseti.Add(new(model.SysUser), "platform")
	mseti.Add(new(model.SysUserBinding))
	mseti.Add(new(model.SysUserTag))
	mseti.Add(new(model.SysUserTemplate))
	mseti.Add(new(model.SysUserTemplateDetail))
	return nil
}

// SyncCtr defined
func SyncCtr() error {
	SysAppFunRoutes(App)
	SysAreaRoutes(App)
	SysAttachmentRoutes(App)
	SysCasRoutes(App)
	SysClientRoutes(App)
	SysCommentRoutes(App)
	SysDataPermissionRoutes(App)
	DebugRoutes(App)
	SysDingtalkRoutes(App)
	SysDomainRoutes(App)
	SysMenuRoutes(App)
	SysNotificationRoutes(App)
	SysOptionsetRoutes(App)
	SysOrgRoutes(App)
	SysPermissionRoutes(App)
	SysRoleRoutes(App)
	SysRoleMenuRoutes(App)
	SysScheduleRoutes(App)
	SysScheduleHistoryRoutes(App)
	SysSchedulingRoutes(App)
	SysSettingRoutes(App)
	SysTableRoutes(App)
	SysTableColumnRoutes(App)
	SysTagRoutes(App)
	SysTagGroupRoutes(App)
	SysTrackerRoutes(App)
	SysUserRoutes(App)
	SysUserTemplateRoutes(App)
	SysUserTemplateDetailRoutes(App)
	SysWechatRoutes(App)
	SysWorkerRoutes(App)
	return nil
}

// SyncService defined
func SyncService() error {
	ClientSrvService(App)
	DomainSrvService(App)
	UserSrvService(App)
	return nil
}

// Executor defined
var Executor = util.NewExecutor(SyncModel, SyncCtr, SyncService)

func init() {
	if err := Executor.Execute(); err != nil {
		panic(err)
	}
}
