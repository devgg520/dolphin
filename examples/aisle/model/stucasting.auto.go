// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Stucasting defined
type Stucasting struct {
	//
	StucastingId null.Int `xorm:"int(11) pk notnull autoincr 'stucasting_id'" json:"stucasting_id" xml:"stucasting_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	CastName null.String `xorm:"varchar(50) 'cast_name'" json:"cast_name" xml:"cast_name"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined Stucasting
func (m *Stucasting) TableName() string {
	return "stucasting"
}
