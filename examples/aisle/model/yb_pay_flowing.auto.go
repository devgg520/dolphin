// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
)

// YbPayFlowing defined
type YbPayFlowing struct {
	//
	YPFId null.Int `xorm:"int(11) pk notnull autoincr 'y_p_f_id'" json:"y_p_f_id" xml:"y_p_f_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" xml:"order_id"`
	//
	PayMoney decimal.Decimal `xorm:"decimal(50,3) 'pay_money'" json:"pay_money" xml:"pay_money"`
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
	YbFlowNumber null.String `xorm:"varchar(1000) 'yb_flow_number'" json:"yb_flow_number" xml:"yb_flow_number"`
	//
	SuccessFailure null.Int `xorm:"int(11) 'success_failure'" json:"success_failure" xml:"success_failure"`
	//
	SourceType null.Int `xorm:"int(11) 'source_type'" json:"source_type" xml:"source_type"`
	//
	Posmenu null.Int `xorm:"int(11) 'posmenu'" json:"posmenu" xml:"posmenu"`
	//
	PayTime null.Time `xorm:"datetime 'pay_time'" json:"pay_time" xml:"pay_time"`
	//
	LoginNum null.String `xorm:"varchar(50) 'login_num'" json:"login_num" xml:"login_num"`
	//
	PosNum null.String `xorm:"varchar(20) 'pos_num'" json:"pos_num" xml:"pos_num"`
	//
	Posrequestid null.String `xorm:"varchar(6) 'posrequestid'" json:"posrequestid" xml:"posrequestid"`
	//
	Referno null.String `xorm:"varchar(50) 'referno'" json:"referno" xml:"referno"`
	//
	Chequeno null.String `xorm:"varchar(50) 'chequeno'" json:"chequeno" xml:"chequeno"`
	//
	Bankcardno null.String `xorm:"varchar(50) 'bankcardno'" json:"bankcardno" xml:"bankcardno"`
	//
	Bankcardname null.String `xorm:"varchar(50) 'bankcardname'" json:"bankcardname" xml:"bankcardname"`
	//
	Bankcardtype null.String `xorm:"varchar(50) 'bankcardtype'" json:"bankcardtype" xml:"bankcardtype"`
	//
	Bankorderno null.String `xorm:"varchar(50) 'bankorderno'" json:"bankorderno" xml:"bankorderno"`
	//
	Yeepayorderno null.String `xorm:"varchar(50) 'yeepayorderno'" json:"yeepayorderno" xml:"yeepayorderno"`
	//
	Customerno null.String `xorm:"varchar(50) 'customerno'" json:"customerno" xml:"customerno"`
	//
	Paytypecode null.Int `xorm:"int(11) 'paytypecode'" json:"paytypecode" xml:"paytypecode"`
	//
	PayState null.Int `xorm:"int(11) 'pay_state'" json:"pay_state" xml:"pay_state"`
	//
	PkSchId null.Int `xorm:"int(11) 'pk_sch_id'" json:"pk_sch_id" xml:"pk_sch_id"`
	//
	YeePayNum null.String `xorm:"varchar(20) 'yee_pay_num'" json:"yee_pay_num" xml:"yee_pay_num"`
	//
	Amount decimal.Decimal `xorm:"decimal(50,3) 'amount'" json:"amount" xml:"amount"`
	//
	AccountState null.Int `xorm:"int(11) 'account_state'" json:"account_state" xml:"account_state"`
	//
	ReconciliationsState null.Int `xorm:"int(11) 'reconciliations_state'" json:"reconciliations_state" xml:"reconciliations_state"`
	//
	ErrorMsg null.String `xorm:"varchar(500) 'error_msg'" json:"error_msg" xml:"error_msg"`
	//
	PkYrrId null.Int `xorm:"int(11) 'pk_yrr_id'" json:"pk_yrr_id" xml:"pk_yrr_id"`
	//
	ServiceCharge decimal.Decimal `xorm:"decimal(50,3) 'service_charge'" json:"service_charge" xml:"service_charge"`
	//
	GetMoney decimal.Decimal `xorm:"decimal(50,3) 'get_money'" json:"get_money" xml:"get_money"`
	//
	UpdateStuybzfDesc null.String `xorm:"varchar(1000) 'update_stuybzf_desc'" json:"update_stuybzf_desc" xml:"update_stuybzf_desc"`
	//
	YpfId null.Int `xorm:"int(11) 'ypf_id'" json:"ypf_id" xml:"ypf_id"`
}

// TableName table name of defined YbPayFlowing
func (m *YbPayFlowing) TableName() string {
	return "yb_pay_flowing"
}
