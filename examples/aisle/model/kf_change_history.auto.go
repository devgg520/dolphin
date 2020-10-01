// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// KfChangeHistory defined
type KfChangeHistory struct {
	//
	KCHId null.Int `xorm:"int(11) pk notnull autoincr 'k_c_h_id'" json:"k_c_h_id" xml:"k_c_h_id"`
	//
	PkSc null.Int `xorm:"int(11) 'pk_sc'" json:"pk_sc" xml:"pk_sc"`
	//
	PkOldKf null.Int `xorm:"int(11) 'pk_old_kf'" json:"pk_old_kf" xml:"pk_old_kf"`
	//
	PkNewKf null.Int `xorm:"int(11) 'pk_new_kf'" json:"pk_new_kf" xml:"pk_new_kf"`
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

// TableName table name of defined KfChangeHistory
func (m *KfChangeHistory) TableName() string {
	return "kf_change_history"
}
