// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	"github.com/gin-gonic/gin"

	platformApp "github.com/2637309949/dolphin/cli/platform/app"
	"github.com/xormplus/xorm"
)

// Engine defined parse app engine
type Engine struct {
	*platformApp.Engine
}

// Context defined http handle hook context
type Context struct {
	*platformApp.Context
}

// RouterGroup defines struct that extend from gin.RouterGroup
type RouterGroup struct {
	*platformApp.RouterGroup
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// HandlerFunc convert to platformApp.HandlerFunc
// rewrite if you need
func (h HandlerFunc) HandlerFunc(e *platformApp.Engine) platformApp.HandlerFunc {
	pgc := platformApp.HandlerFunc(func(ctx *platformApp.Context) {
		h(&Context{Context: ctx})
	})
	pgc.HandlerFunc(e)
	return pgc
}

// BuildEngine build engine
func BuildEngine(build func(*Engine)) func(*platformApp.Engine) {
	return func(e *platformApp.Engine) {
		engine.Engine = e
		build(engine)
	}
}

// Query defined
func (e *Engine) Query(ctx *Context) *platformApp.Query {
	return e.Engine.Query(ctx.Context)
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{RouterGroup: e.Engine.Group(relativePath, handlers...)}
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	var newHandlers []platformApp.HandlerFunc
	for _, h := range handlers {
		newHandlers = append(newHandlers, h.HandlerFunc(rg.Engine))
	}
	return rg.RouterGroup.Handle(httpMethod, relativePath, newHandlers...)
}

// PageSearch Rewrite if you need
func (e *Engine) PageSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
	return e.Engine.PageSearch(db, controller, api, table, q)
}

// Run booting system
func (e *Engine) Run() {
	e.Engine.Run()
}

// Engine instance
var engine = &Engine{}
