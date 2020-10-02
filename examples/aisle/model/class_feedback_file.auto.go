// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassFeedbackFile defined
type ClassFeedbackFile struct {
	//
	CFFId null.Int `xorm:"int(11) pk notnull autoincr 'c_f_f_id'" json:"c_f_f_id" xml:"c_f_f_id"`
	//
	CfId null.Int `xorm:"int(11) 'cf_id'" json:"cf_id" xml:"cf_id"`
	//
	Filed null.Int `xorm:"int(11) 'filed'" json:"filed" xml:"filed"`
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

// TableName table name of defined ClassFeedbackFile
func (m *ClassFeedbackFile) TableName() string {
	return "class_feedback_file"
}