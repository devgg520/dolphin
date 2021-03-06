// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuAllotTmk defined
type StuAllotTmk struct {
	//
	SATId null.Int `xorm:"int(11) pk notnull autoincr 's_a_t_id'" json:"s_a_t_id" xml:"s_a_t_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	TmkUserid null.Int `xorm:"int(11) 'tmk_userid'" json:"tmk_userid" xml:"tmk_userid"`
	//
	AllotState null.Int `xorm:"int(11) 'allot_state'" json:"allot_state" xml:"allot_state"`
	//
	AllotDate null.Time `xorm:"datetime 'allot_date'" json:"allot_date" xml:"allot_date"`
	//
	QxAllotDate null.Time `xorm:"datetime 'qx_allot_date'" json:"qx_allot_date" xml:"qx_allot_date"`
	//
	HisTmkUserid null.Int `xorm:"int(11) 'his_tmk_userid'" json:"his_tmk_userid" xml:"his_tmk_userid"`
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
	AllotDesc null.String `xorm:"varchar(5000) 'allot_desc'" json:"allot_desc" xml:"allot_desc"`
	//
	IfPl null.Int `xorm:"int(11) 'if_pl'" json:"if_pl" xml:"if_pl"`
}

// TableName table name of defined StuAllotTmk
func (m *StuAllotTmk) TableName() string {
	return "stu_allot_tmk"
}
