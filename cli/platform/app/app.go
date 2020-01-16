// Code generated by dol build. Only Generate by tools if not existed.
// source: app.go

package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/2637309949/dolphin/cli/packages/logrus"
	"github.com/2637309949/dolphin/cli/packages/viper"

	// github.com/2637309949/dolphin/cli/platform/conf
	_ "github.com/2637309949/dolphin/cli/platform/conf"
	"github.com/2637309949/dolphin/srv"
	"github.com/2637309949/dolphin/srv/cli"
)

// HTTPServer defined http.Server
var HTTPServer *http.Server

// InitServer defined HTTPServer
func InitServer() {
	if HTTPServer == nil {
		HTTPServer = &http.Server{Addr: fmt.Sprintf(":%v", viper.GetString("http.port"))}
	}
}

// RPCListener defined net
var RPCListener net.Listener

// InitRPCListener defined RPCListener
func InitRPCListener() {
	if RPCListener == nil {
		grpc, err := net.Listen("tcp", fmt.Sprintf(":%v", viper.GetString("grpc.port")))
		if err != nil {
			logrus.Fatal(err)
		}
		RPCListener = grpc
	}
}

// OnStart defined OnStart
func OnStart(e *Engine) func(context.Context) error {
	return func(context.Context) error {
		go func() {
			logrus.Infof("http listen on port:%v", viper.GetString("http.port"))
			HTTPServer.Handler = e.Gin
			if err := HTTPServer.ListenAndServe(); err != nil {
				logrus.Fatal(err)
			}
		}()
		go func() {
			logrus.Infof("grpc listen on port:%v", viper.GetString("grpc.port"))
			if err := e.GRPC.Serve(RPCListener); err != nil {
				logrus.Fatal(err)
			}
		}()
		return nil
	}
}

// OnStop defined OnStop
func OnStop(e *Engine) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		if err := HTTPServer.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
			return err
		}
		if err := RPCListener.Close(); err != nil {
			logrus.Fatal(err)
			return err
		}
		return nil
	}
}

// NewLifeHook create lifecycle hook
func NewLifeHook(e *Engine) srv.Hook {
	InitServer()
	InitRPCListener()
	return srv.Hook{
		OnStart: OnStart(e),
		OnStop:  OnStop(e),
	}
}

func init() {
	authServerURL = viper.GetString("oauth.server")
	oa2cfg.ClientID = viper.GetString("oauth.id")
	oa2cfg.ClientSecret = viper.GetString("oauth.secret")
	oa2cfg.RedirectURL = fmt.Sprintf("%v/api/sys/oauth2/oauth2", viper.GetString("oauth.cli"))
	oa2cfg.Endpoint.AuthURL = authServerURL + "/api/sys/oauth2/authorize"
	oa2cfg.Endpoint.TokenURL = authServerURL + "/api/sys/oauth2/token"
	_ = cli.Provider(func(lc srv.Lifecycle) *Engine {
		var engine = NewEngine()
		lc.Append(NewLifeHook(engine))
		return engine
	})
}
