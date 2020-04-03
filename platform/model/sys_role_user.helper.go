// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// DefaultRoleUser default admin
var DefaultRoleUser = SysRoleUser{
	ID:         null.StringFrom("03eda436-2772-48da-86d8-97b2bd80e391"),
	UserId:     DefaultAdmin.ID,
	RoleId:     DefaultRole.ID,
	CreateBy:   DefaultAdmin.ID,
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   DefaultAdmin.ID,
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}

// InitSysData defined inital system data
func (m *SysRoleUser) InitSysData(s *xorm.Session) {
	if ct, err := s.Where("id=?", DefaultRoleUser.ID).Count(new(SysRoleUser)); ct == 0 || err != nil {
		if err != nil {
			s.Rollback()
			panic(err)
		}
		if _, err := s.InsertOne(&DefaultRoleUser); err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
