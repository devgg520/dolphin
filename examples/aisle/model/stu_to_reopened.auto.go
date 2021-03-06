// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuToReopened defined
type StuToReopened struct {
	//
	STRId null.Int `xorm:"int(11) pk notnull autoincr 's_t_r_id'" json:"s_t_r_id" xml:"s_t_r_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	StrRemark null.String `xorm:"varchar(2000) 'str_remark'" json:"str_remark" xml:"str_remark"`
	//
	StrDate null.Time `xorm:"datetime 'str_date'" json:"str_date" xml:"str_date"`
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
	StrBussType null.Int `xorm:"int(11) 'str_buss_type'" json:"str_buss_type" xml:"str_buss_type"`
}

// TableName table name of defined StuToReopened
func (m *StuToReopened) TableName() string {
	return "stu_to_reopened"
}
