// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/2637309949/dolphin/packages/cache"

	"github.com/2637309949/dolphin/packages/cache/persistence"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2"
	"github.com/2637309949/dolphin/packages/oauth2/store"
	"github.com/2637309949/dolphin/packages/redis"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/http"
)

// Manager Engine management interface
type Manager interface {
	MSet() MSeti
	GetBusinessDBSet() map[string]*xorm.Engine
	GetBusinessDB(string) *xorm.Engine
	AddBusinessDB(string, *xorm.Engine)
	GetTokenStore() oauth2.TokenStore
}

// DefaultManager defined
type DefaultManager struct {
	BusinessDBSet map[string]*xorm.Engine
	MSeti         MSeti
}

// MSet defined
func (d *DefaultManager) MSet() MSeti {
	return d.MSeti
}

// GetBusinessDB defined
func (d *DefaultManager) GetBusinessDB(domain string) *xorm.Engine {
	return d.BusinessDBSet[domain]
}

// GetBusinessDBSet defined
func (d *DefaultManager) GetBusinessDBSet() map[string]*xorm.Engine {
	return d.BusinessDBSet
}

// AddBusinessDB defined
func (d *DefaultManager) AddBusinessDB(domain string, db *xorm.Engine) {
	if d.BusinessDBSet[domain] != nil {
		panic(fmt.Errorf("domain %v has exited", domain))
	}
	d.BusinessDBSet[domain] = db
}

// GetTokenStore defined
func (d *DefaultManager) GetTokenStore() oauth2.TokenStore {
	if RedisClient != nil {
		return store.NewRedisStoreWithCli(RedisClient, TokenkeyNamespace)
	}
	memo, _ := store.NewMemoryTokenStore()
	return memo
}

// NewDefaultManager defined
func NewDefaultManager() Manager {
	mg := &DefaultManager{}
	mg.BusinessDBSet = map[string]*xorm.Engine{}
	mg.MSeti = &MSet{m: map[string][]interface{}{}}
	return mg
}

// Cache middles
func Cache(time time.Duration) func(ctx *Context) {
	return func(ctx *Context) {
		cache.CachePage(CacheStore, time)(ctx.Context)
	}
}

// RedisClient defined
var RedisClient *redis.Client

// CacheStore defined
var CacheStore persistence.CacheStore

func init() {
	if uri := util.EnsureLeft(http.Parse(viper.GetString("rd.dataSource"))).(*http.URI); uri.Laddr != "" {
		RedisClient := redis.NewClient(&redis.Options{
			Addr:     uri.Laddr,
			Password: uri.Passwd,
			DB:       util.EnsureLeft(strconv.Atoi(uri.DbName)).(int),
		})
		if _, err := RedisClient.Ping().Result(); err != nil {
			logrus.Warnf("Redis:%v connect failed", viper.GetString("rd.dataSource"))
			RedisClient = nil
		}
	}
	CacheStore = persistence.NewInMemoryStore(60 * time.Second)
}
