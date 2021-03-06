// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuGwcGef defined
type StuGwcGef struct {
	//
	SGGId null.Int `xorm:"int(11) pk notnull autoincr 's_g_g_id'" json:"s_g_g_id" xml:"s_g_g_id"`
	//
	StuGwcId null.Int `xorm:"int(11) 'stu_gwc_id'" json:"stu_gwc_id" xml:"stu_gwc_id"`
	//
	GefId null.Int `xorm:"int(11) 'gef_id'" json:"gef_id" xml:"gef_id"`
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

// TableName table name of defined StuGwcGef
func (m *StuGwcGef) TableName() string {
	return "stu_gwc_gef"
}
