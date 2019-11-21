// Code generated by dol build. Only Generate by tools if not existed, your can rewrite platform.App default action
// source: app.go

package app

import (
	pApp "github.com/2637309949/dolphin/cli/platform/app"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type (
	// Engine defined parse app engine
	Engine struct {
		*pApp.Engine
	}
	// Context defined http handle hook context
	Context struct {
		*pApp.Context
	}
	// RouterGroup defines struct that extend from gin.RouterGroup
	RouterGroup struct {
		*pApp.RouterGroup
	}
	// HandlerFunc defines the handler used by gin middleware as return value.
	HandlerFunc func(*Context)
)

// BuildEngine build engine
func BuildEngine(build func(*Engine)) func(*pApp.Engine) {
	return func(base *pApp.Engine) {
		engine.Engine = base
		build(engine)
	}
}

// Auth middles
func (e *Engine) Auth(h func(*Context)) func(*Context) {
	return func(ctx *Context) {
		e.Engine.Auth(func(*pApp.Context) {
			h(ctx)
		})(ctx.Context)
	}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{RouterGroup: e.Engine.Group(relativePath, handlers...)}
}

// HandlerFunc convert to pApp.HandlerFunc
func (hf HandlerFunc) HandlerFunc(e *pApp.Engine) (phf pApp.HandlerFunc) {
	phf = pApp.HandlerFunc(func(base *pApp.Context) {
		hf(&Context{Context: base})
	})
	phf.HandlerFunc(e)
	return
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) pApp.HandlerFunc {
		return h.HandlerFunc(rg.Engine)
	}).([]pApp.HandlerFunc)...)
}

// Engine instance
var engine = &Engine{}
