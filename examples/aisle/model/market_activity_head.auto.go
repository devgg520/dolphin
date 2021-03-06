// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketActivityHead defined
type MarketActivityHead struct {
	//
	MAHId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_h_id'" json:"m_a_h_id" xml:"m_a_h_id"`
	//
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" xml:"ma_id"`
	//
	HeadId null.Int `xorm:"int(11) 'head_id'" json:"head_id" xml:"head_id"`
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

// TableName table name of defined MarketActivityHead
func (m *MarketActivityHead) TableName() string {
	return "market_activity_head"
}
