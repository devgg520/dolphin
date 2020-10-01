// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuBfJl defined
type StuBfJl struct {
	//
	StuBfJlId null.Int `xorm:"int(11) pk notnull autoincr 'stu_bf_jl_id'" json:"stu_bf_jl_id" xml:"stu_bf_jl_id"`
	//
	LbfId null.Int `xorm:"int(11) 'lbf_id'" json:"lbf_id" xml:"lbf_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	BfSc null.String `xorm:"varchar(500) 'bf_sc'" json:"bf_sc" xml:"bf_sc"`
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

// TableName table name of defined StuBfJl
func (m *StuBfJl) TableName() string {
	return "stu_bf_jl"
}
