// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// RefundStuClassType defined
type RefundStuClassType struct {
	//
	RSCTId null.Int `xorm:"int(11) pk notnull autoincr 'r_s_c_t_id'" json:"r_s_c_t_id" xml:"r_s_c_t_id"`
	//
	RefId null.Int `xorm:"int(11) 'ref_id'" json:"ref_id" xml:"ref_id"`
	//
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" xml:"sct_id"`
	//
	Money null.Float `xorm:"float(10,2) 'money'" json:"money" xml:"money"`
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
	RefundHour null.Float `xorm:"float(10,2) 'refund_hour'" json:"refund_hour" xml:"refund_hour"`
	//
	RefUnitprice null.Float `xorm:"float(10,2) 'ref_unitprice'" json:"ref_unitprice" xml:"ref_unitprice"`
	//
	RstInvoiceFee null.Float `xorm:"float(50,2) 'rst_invoice_fee'" json:"rst_invoice_fee" xml:"rst_invoice_fee"`
	//
	RstServiceFee null.Float `xorm:"float(50,2) 'rst_service_fee'" json:"rst_service_fee" xml:"rst_service_fee"`
	//
	CheckState null.Int `xorm:"int(11) default(54) 'check_state'" json:"check_state" xml:"check_state"`
	//
	RefWay null.Int `xorm:"int(11) 'ref_way'" json:"ref_way" xml:"ref_way"`
	//
	RefType null.Int `xorm:"int(11) 'ref_type'" json:"ref_type" xml:"ref_type"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" xml:"order_id"`
}

// TableName table name of defined RefundStuClassType
func (m *RefundStuClassType) TableName() string {
	return "refund_stu_class_type"
}
