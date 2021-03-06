// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
)

// YeepayReconciliationsRecord defined
type YeepayReconciliationsRecord struct {
	//
	YRRId null.Int `xorm:"int(11) pk notnull autoincr 'y_r_r_id'" json:"y_r_r_id" xml:"y_r_r_id"`
	//
	ConsumerName null.String `xorm:"varchar(100) 'consumer_name'" json:"consumer_name" xml:"consumer_name"`
	//
	ConsumerNum null.String `xorm:"varchar(100) 'consumer_num'" json:"consumer_num" xml:"consumer_num"`
	//
	NetConsumerNum null.String `xorm:"varchar(100) 'net_consumer_num'" json:"net_consumer_num" xml:"net_consumer_num"`
	//
	NetConsumerName null.String `xorm:"varchar(100) 'net_consumer_name'" json:"net_consumer_name" xml:"net_consumer_name"`
	//
	OrderNum null.String `xorm:"varchar(100) 'order_num'" json:"order_num" xml:"order_num"`
	//
	Yeepayorderno null.String `xorm:"varchar(100) 'yeepayorderno'" json:"yeepayorderno" xml:"yeepayorderno"`
	//
	TransactionType null.String `xorm:"varchar(100) 'transaction_type'" json:"transaction_type" xml:"transaction_type"`
	//
	Amount decimal.Decimal `xorm:"decimal(11,2) 'amount'" json:"amount" xml:"amount"`
	//
	ServiceCharge decimal.Decimal `xorm:"decimal(11,2) 'service_charge'" json:"service_charge" xml:"service_charge"`
	//
	GetMoney decimal.Decimal `xorm:"decimal(11,2) 'get_money'" json:"get_money" xml:"get_money"`
	//
	PayState null.String `xorm:"varchar(10) 'pay_state'" json:"pay_state" xml:"pay_state"`
	//
	RefundNum null.String `xorm:"varchar(100) 'refund_num'" json:"refund_num" xml:"refund_num"`
	//
	ConsumerRefundNum null.String `xorm:"varchar(100) 'consumer_refund_num'" json:"consumer_refund_num" xml:"consumer_refund_num"`
	//
	RefundState null.String `xorm:"varchar(10) 'refund_state'" json:"refund_state" xml:"refund_state"`
	//
	RefundMoney decimal.Decimal `xorm:"decimal(11,2) 'refund_money'" json:"refund_money" xml:"refund_money"`
	//
	PosTerminalNum null.String `xorm:"varchar(100) 'pos_terminal_num'" json:"pos_terminal_num" xml:"pos_terminal_num"`
	//
	PosNum null.String `xorm:"varchar(100) 'pos_num'" json:"pos_num" xml:"pos_num"`
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
	PosSerialNum null.String `xorm:"varchar(50) 'pos_serial_num'" json:"pos_serial_num" xml:"pos_serial_num"`
	//
	ChargeBatchNum null.String `xorm:"varchar(50) 'charge_batch_num'" json:"charge_batch_num" xml:"charge_batch_num"`
	//
	Referno null.String `xorm:"varchar(50) 'referno'" json:"referno" xml:"referno"`
	//
	AuthorizationNum null.String `xorm:"varchar(50) 'authorization_num'" json:"authorization_num" xml:"authorization_num"`
	//
	ChargeRequestDate null.Time `xorm:"datetime 'charge_request_date'" json:"charge_request_date" xml:"charge_request_date"`
	//
	ChargeSuccessDate null.Time `xorm:"datetime 'charge_success_date'" json:"charge_success_date" xml:"charge_success_date"`
	//
	RefundRequestDate null.Time `xorm:"datetime 'refund_request_date'" json:"refund_request_date" xml:"refund_request_date"`
	//
	Bankcardno null.String `xorm:"varchar(20) 'bankcardno'" json:"bankcardno" xml:"bankcardno"`
	//
	Bankcardtype null.String `xorm:"varchar(10) 'bankcardtype'" json:"bankcardtype" xml:"bankcardtype"`
	//
	Bankcardname null.String `xorm:"varchar(50) 'bankcardname'" json:"bankcardname" xml:"bankcardname"`
	//
	PayType null.String `xorm:"varchar(10) 'pay_type'" json:"pay_type" xml:"pay_type"`
	//
	ReconciliationsState null.Int `xorm:"int(11) 'reconciliations_state'" json:"reconciliations_state" xml:"reconciliations_state"`
	//
	ErrorMsg null.String `xorm:"varchar(500) 'error_msg'" json:"error_msg" xml:"error_msg"`
	//
	PkYpfId null.Int `xorm:"int(11) 'pk_ypf_id'" json:"pk_ypf_id" xml:"pk_ypf_id"`
	//
	AccountState null.Int `xorm:"int(11) 'account_state'" json:"account_state" xml:"account_state"`
	//
	IfOwnRecon null.Int `xorm:"int(11) 'if_own_recon'" json:"if_own_recon" xml:"if_own_recon"`
	//
	IfPerfect null.Int `xorm:"int(11) 'if_perfect'" json:"if_perfect" xml:"if_perfect"`
}

// TableName table name of defined YeepayReconciliationsRecord
func (m *YeepayReconciliationsRecord) TableName() string {
	return "yeepay_reconciliations_record"
}
