// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CcsHomeworkCheckFile defined
type CcsHomeworkCheckFile struct {
	//
	CHCFId null.Int `xorm:"int(11) pk notnull autoincr 'c_h_c_f_id'" json:"c_h_c_f_id" xml:"c_h_c_f_id"`
	//
	CctId null.Int `xorm:"int(11) 'cct_id'" json:"cct_id" xml:"cct_id"`
	//
	File null.Int `xorm:"int(11) 'file'" json:"file" xml:"file"`
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
	ScsId null.Int `xorm:"int(11) 'scs_id'" json:"scs_id" xml:"scs_id"`
}

// TableName table name of defined CcsHomeworkCheckFile
func (m *CcsHomeworkCheckFile) TableName() string {
	return "ccs_homework_check_file"
}
