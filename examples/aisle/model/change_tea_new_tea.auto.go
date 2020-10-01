// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ChangeTeaNewTea defined
type ChangeTeaNewTea struct {
	//
	CTNTId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_n_t_id'" json:"c_t_n_t_id" xml:"c_t_n_t_id"`
	//
	PkChangeTea null.Int `xorm:"int(11) 'pk_change_tea'" json:"pk_change_tea" xml:"pk_change_tea"`
	//
	PkNewTea null.Int `xorm:"int(11) 'pk_new_tea'" json:"pk_new_tea" xml:"pk_new_tea"`
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

// TableName table name of defined ChangeTeaNewTea
func (m *ChangeTeaNewTea) TableName() string {
	return "change_tea_new_tea"
}
