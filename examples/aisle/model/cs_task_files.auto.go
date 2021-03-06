// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CsTaskFiles defined
type CsTaskFiles struct {
	//
	CTFId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_f_id'" json:"c_t_f_id" xml:"c_t_f_id"`
	//
	CsTaskId null.Int `xorm:"int(11) 'cs_task_id'" json:"cs_task_id" xml:"cs_task_id"`
	//
	FileId null.Int `xorm:"int(11) 'file_id'" json:"file_id" xml:"file_id"`
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

// TableName table name of defined CsTaskFiles
func (m *CsTaskFiles) TableName() string {
	return "cs_task_files"
}
