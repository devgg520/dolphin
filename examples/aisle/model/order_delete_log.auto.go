// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// OrderDeleteLog defined
type OrderDeleteLog struct {
	//
	ODLId null.Int `xorm:"int(11) pk notnull autoincr 'o_d_l_id'" json:"o_d_l_id" xml:"o_d_l_id"`
	//
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" xml:"order_id"`
	//
	DeleteTime null.Time `xorm:"datetime 'delete_time'" json:"delete_time" xml:"delete_time"`
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
	OfNumber null.String `xorm:"varchar(100) 'of_number'" json:"of_number" xml:"of_number"`
	//
	StuPhone null.String `xorm:"varchar(20) 'stu_phone'" json:"stu_phone" xml:"stu_phone"`
	//
	StuPhone2 null.String `xorm:"varchar(20) 'stu_phone2'" json:"stu_phone2" xml:"stu_phone2"`
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" xml:"sch_id"`
	//
	OfState null.Int `xorm:"int(11) 'of_state'" json:"of_state" xml:"of_state"`
}

// TableName table name of defined OrderDeleteLog
func (m *OrderDeleteLog) TableName() string {
	return "order_delete_log"
}
