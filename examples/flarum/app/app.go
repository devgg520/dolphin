// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	"sync"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	pApp "github.com/2637309949/dolphin/platform/app"
)

type (
	// Engine defined parse app engine
	Engine struct {
		*pApp.Engine
		pool sync.Pool
	}
	// Context defined http handle hook context
	Context struct {
		*pApp.Context
		engine *Engine
	}
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		*pApp.RouterGroup
		engine *Engine
	}
	// HandlerFunc defines the handler used by gin middleware as return value.
	HandlerFunc func(*Context)
)

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Engine.Group(relativePath, handlers...)}
}

// HandlerFunc convert to pApp.HandlerFunc
func (e *Engine) HandlerFunc(h HandlerFunc) (phf pApp.HandlerFunc) {
	return pApp.HandlerFunc(func(ctx *pApp.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		h(c)
		e.pool.Put(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) []gin.IRoutes {
	rh := rg.RouterGroup.Handle(
		httpMethod,
		relativePath,
		funk.Map(handlers, func(h HandlerFunc) pApp.HandlerFunc {
			return rg.engine.HandlerFunc(h)
		}).([]pApp.HandlerFunc)...)
	return rh
}

// Auth middles
func Auth(ctx *Context) {
	pApp.Auth(ctx.Context)
}

// Roles middles
func Roles(roles ...string) func(ctx *Context) {
	return func(ctx *Context) {
		pApp.Roles(roles...)(ctx.Context)
	}
}

// Cache middles
func Cache(time time.Duration) func(ctx *Context) {
	return func(ctx *Context) {
		pApp.Cache(time)(ctx.Context)
	}
}

// buildEngine defined init engine you can custom engine
// if you need
func buildEngine() *Engine {
	e := &Engine{Engine: pApp.App}
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}

// Run app
func Run() {
	App.Run()
}

// App instance
var App = buildEngine()
