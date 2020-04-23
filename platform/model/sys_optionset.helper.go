package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// InitSysData defined inital system data
func (m *SysOptionset) InitSysData(s *xorm.Session) {
	options := []SysOptionset{
		SysOptionset{
			ID:         null.StringFrom("c6a7bt50-7dad-31d1-81b5-10c34fd460e2"),
			Name:       null.StringFrom("域名启用状态"),
			Code:       null.StringFrom("sys_domain_status"),
			Value:      null.StringFrom(`[{"text":"禁用","value":0},{"text":"正常","value":1}]`),
			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		SysOptionset{
			ID:         null.StringFrom("b6a7ct5q-2dad-31d1-81b5-10c34fd460c3"),
			Name:       null.StringFrom("域名认证模式"),
			Code:       null.StringFrom("sys_domain_auth_mode"),
			Value:      null.StringFrom(`[{"text":"集成登录","value":0},{"text":"单点登录","value":1}]`),
			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
		SysOptionset{
			ID:         null.StringFrom("c6a7bt50-7dad-31d1-81b5-10c34fd460c2"),
			Name:       null.StringFrom("域名数据库标志"),
			Code:       null.StringFrom("sys_domain_sync_flag"),
			Value:      null.StringFrom(`[{"text":"未同步","value":0},{"text":"已同步","value":1}]`),
			CreateBy:   DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			UpdateBy:   DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			DelFlag:    null.IntFrom(0),
		},
	}
	for _, option := range options {
		if ct, err := s.Where("id=?", option.ID.String).Count(new(SysOptionset)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			if _, err := s.InsertOne(&option); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}