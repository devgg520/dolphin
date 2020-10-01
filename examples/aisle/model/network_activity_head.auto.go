// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// NetworkActivityHead defined
type NetworkActivityHead struct {
	//
	NAHId null.Int `xorm:"int(11) pk notnull autoincr 'n_a_h_id'" json:"n_a_h_id" xml:"n_a_h_id"`
	//
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" xml:"na_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined NetworkActivityHead
func (m *NetworkActivityHead) TableName() string {
	return "network_activity_head"
}
