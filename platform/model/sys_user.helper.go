package model

import (
	"fmt"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/util"
	"golang.org/x/crypto/scrypt"
)

// DefaultAdmin default admin
var DefaultAdmin = SysUser{
	ID:         null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	Password:   null.StringFrom("admin"),
	Name:       null.StringFrom("admin"),
	Nickname:   null.StringFrom("admin"),
	Avatar:     null.StringFrom("http://pic.616pic.com/ys_bnew_img/00/06/27/TWk2P5YJ5k.jpg"),
	Status:     null.IntFrom(1),
	Domain:     null.StringFrom("localhost"),
	OrgId:      null.StringFrom("c637bt50-7dad-31d1-81b5-10c34fd460e1"),
	CreateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	CreateTime: null.TimeFrom(time.Now()),
	UpdateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	UpdateTime: null.TimeFrom(time.Now()),
	DelFlag:    null.IntFrom(0),
}

// SetPassword Method to set salt and hash the password for a user
func (m *SysUser) SetPassword(password string) {
	b := util.RandString(16, util.RandNumChar)
	m.Salt = null.StringFrom(b)
	dk, err := scrypt.Key([]byte(password), []byte(m.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	m.Password = null.StringFrom(fmt.Sprintf("%x", dk))
}

// ValidPassword Method to check the entered password is correct or not
func (m *SysUser) ValidPassword(password string) bool {
	dk, err := scrypt.Key([]byte(password), []byte(m.Salt.String), 512, 8, 1, 64)
	if err != nil {
		panic(err)
	}
	return m.Password.String == fmt.Sprintf("%x", dk)
}

// InitSysData defined inital system data
func (m *SysUser) InitSysData(s *xorm.Session) {
	if ct, err := s.Where("id=?", DefaultAdmin.ID.String).Count(new(SysUser)); ct == 0 || err != nil {
		if err != nil {
			s.Rollback()
			panic(err)
		}
		DefaultAdmin.SetPassword(DefaultAdmin.Password.String)
		if _, err := s.InsertOne(&DefaultAdmin); err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
