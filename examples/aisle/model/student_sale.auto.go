// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentSale defined
type StudentSale struct {
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	SSId null.Int `xorm:"int(11) pk notnull autoincr 's_s_id'" json:"s_s_id" xml:"s_s_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	SsState null.Int `xorm:"int(11) 'ss_state'" json:"ss_state" xml:"ss_state"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	AllotSourse null.Int `xorm:"int(11) 'allot_sourse'" json:"allot_sourse" xml:"allot_sourse"`
	//
	AllotDate null.Time `xorm:"datetime 'allot_date'" json:"allot_date" xml:"allot_date"`
	//
	CancelDate null.Time `xorm:"datetime 'cancel_date'" json:"cancel_date" xml:"cancel_date"`
	//
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" xml:"remark"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" xml:"buss_type"`
	//
	YesornoBatchAllocation null.Int `xorm:"int(11) 'yesorno_batch_allocation'" json:"yesorno_batch_allocation" xml:"yesorno_batch_allocation"`
	//
	History null.Int `xorm:"int(11) 'history'" json:"history" xml:"history"`
}

// TableName table name of defined StudentSale
func (m *StudentSale) TableName() string {
	return "student_sale"
}
