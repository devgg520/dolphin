// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuFailToLive defined
type StuFailToLive struct {
	//
	SFTLId null.Int `xorm:"int(11) pk notnull autoincr 's_f_t_l_id'" json:"s_f_t_l_id" xml:"s_f_t_l_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	SftlRemark null.String `xorm:"varchar(2000) 'sftl_remark'" json:"sftl_remark" xml:"sftl_remark"`
	//
	SftlDate null.Time `xorm:"datetime 'sftl_date'" json:"sftl_date" xml:"sftl_date"`
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
	SftlBussType null.Int `xorm:"int(11) 'sftl_buss_type'" json:"sftl_buss_type" xml:"sftl_buss_type"`
}

// TableName table name of defined StuFailToLive
func (m *StuFailToLive) TableName() string {
	return "stu_fail_to_live"
}
