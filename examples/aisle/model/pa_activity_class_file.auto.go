// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// PaActivityClassFile defined
type PaActivityClassFile struct {
	//
	PACFId null.Int `xorm:"int(11) pk notnull autoincr 'p_a_c_f_id'" json:"p_a_c_f_id" xml:"p_a_c_f_id"`
	//
	FileId null.Int `xorm:"int(11) 'file_id'" json:"file_id" xml:"file_id"`
	//
	PacfDesc null.String `xorm:"varchar(500) 'pacf_desc'" json:"pacf_desc" xml:"pacf_desc"`
	//
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" xml:"par_id"`
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

// TableName table name of defined PaActivityClassFile
func (m *PaActivityClassFile) TableName() string {
	return "pa_activity_class_file"
}
