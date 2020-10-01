// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuExchangeGift defined
type StuExchangeGift struct {
	//
	SEGId null.Int `xorm:"int(11) pk notnull autoincr 's_e_g_id'" json:"s_e_g_id" xml:"s_e_g_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	SegRemark null.String `xorm:"varchar(1000) 'seg_remark'" json:"seg_remark" xml:"seg_remark"`
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
	UseTotalPoint null.Float `xorm:"float(10,2) 'use_total_point'" json:"use_total_point" xml:"use_total_point"`
}

// TableName table name of defined StuExchangeGift
func (m *StuExchangeGift) TableName() string {
	return "stu_exchange_gift"
}
