// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CssDeleteRecord defined
type CssDeleteRecord struct {
	//
	CDRId null.Int `xorm:"int(11) pk notnull autoincr 'c_d_r_id'" json:"c_d_r_id" xml:"c_d_r_id"`
	//
	CssId null.Int `xorm:"int(11) 'css_id'" json:"css_id" xml:"css_id"`
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
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	ScsId null.Int `xorm:"int(11) 'scs_id'" json:"scs_id" xml:"scs_id"`
}

// TableName table name of defined CssDeleteRecord
func (m *CssDeleteRecord) TableName() string {
	return "css_delete_record"
}
