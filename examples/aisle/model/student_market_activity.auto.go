// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentMarketActivity defined
type StudentMarketActivity struct {
	//
	SMAId null.Int `xorm:"int(11) pk notnull autoincr 's_m_a_id'" json:"s_m_a_id" xml:"s_m_a_id"`
	//
	SutId null.Int `xorm:"int(11) 'sut_id'" json:"sut_id" xml:"sut_id"`
	//
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" xml:"ma_id"`
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

// TableName table name of defined StudentMarketActivity
func (m *StudentMarketActivity) TableName() string {
	return "student_market_activity"
}
