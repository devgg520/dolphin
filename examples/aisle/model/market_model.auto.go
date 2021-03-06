// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketModel defined
type MarketModel struct {
	//
	MMId null.Int `xorm:"int(11) pk notnull autoincr 'm_m_id'" json:"m_m_id" xml:"m_m_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	MmName null.String `xorm:"varchar(100) 'mm_name'" json:"mm_name" xml:"mm_name"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	MmType null.Int `xorm:"int(11) 'mm_type'" json:"mm_type" xml:"mm_type"`
	//
	TsId null.Int `xorm:"int(11) 'ts_id'" json:"ts_id" xml:"ts_id"`
	//
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" xml:"school_id"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined MarketModel
func (m *MarketModel) TableName() string {
	return "market_model"
}
