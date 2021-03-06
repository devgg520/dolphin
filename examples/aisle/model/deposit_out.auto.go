// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// DepositOut defined
type DepositOut struct {
	//
	DOId null.Int `xorm:"int(11) pk notnull autoincr 'd_o_id'" json:"d_o_id" xml:"d_o_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	OutMoney null.Float `xorm:"float(11,2) 'out_money'" json:"out_money" xml:"out_money"`
	//
	OutWay null.Int `xorm:"int(11) 'out_way'" json:"out_way" xml:"out_way"`
	//
	OutDate null.Time `xorm:"datetime 'out_date'" json:"out_date" xml:"out_date"`
	//
	Remark null.String `xorm:"varchar(1000) 'remark'" json:"remark" xml:"remark"`
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
	DiId null.Int `xorm:"int(11) 'di_id'" json:"di_id" xml:"di_id"`
	//
	FeeId null.Int `xorm:"int(11) 'fee_id'" json:"fee_id" xml:"fee_id"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" xml:"buss_type"`
	//
	OutType null.Int `xorm:"int(11) 'out_type'" json:"out_type" xml:"out_type"`
	//
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" xml:"check_state"`
	//
	CheckUserId null.Int `xorm:"int(11) 'check_user_id'" json:"check_user_id" xml:"check_user_id"`
	//
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" xml:"check_time"`
	//
	ActualRefundTime null.Time `xorm:"datetime 'actual_refund_time'" json:"actual_refund_time" xml:"actual_refund_time"`
	//
	ZdCheckState null.Int `xorm:"int(11) 'zd_check_state'" json:"zd_check_state" xml:"zd_check_state"`
	//
	CheckDate null.Time `xorm:"datetime 'check_date'" json:"check_date" xml:"check_date"`
	//
	PkFlowSet null.Int `xorm:"int(11) 'pk_flow_set'" json:"pk_flow_set" xml:"pk_flow_set"`
	//
	NowCheckUser null.String `xorm:"varchar(500) 'now_check_user'" json:"now_check_user" xml:"now_check_user"`
	//
	NowFloor null.Int `xorm:"int(11) 'now_floor'" json:"now_floor" xml:"now_floor"`
	//
	TurnFloor null.Int `xorm:"int(11) 'turn_floor'" json:"turn_floor" xml:"turn_floor"`
	//
	HistoryCheckUser null.String `xorm:"varchar(500) 'history_check_user'" json:"history_check_user" xml:"history_check_user"`
	//
	PkCfpId null.Int `xorm:"int(11) 'pk_cfp_id'" json:"pk_cfp_id" xml:"pk_cfp_id"`
	//
	NowCfid null.Int `xorm:"int(11) 'now_cfid'" json:"now_cfid" xml:"now_cfid"`
	//
	NextCfid null.Int `xorm:"int(11) 'next_cfid'" json:"next_cfid" xml:"next_cfid"`
	//
	PkCheckUser null.Int `xorm:"int(11) 'pk_check_user'" json:"pk_check_user" xml:"pk_check_user"`
	//
	JzzmFileid null.Int `xorm:"int(11) 'jzzm_fileid'" json:"jzzm_fileid" xml:"jzzm_fileid"`
	//
	JzbmFileid null.Int `xorm:"int(11) 'jzbm_fileid'" json:"jzbm_fileid" xml:"jzbm_fileid"`
	//
	HmName null.String `xorm:"varchar(100) 'hm_name'" json:"hm_name" xml:"hm_name"`
	//
	BankName null.String `xorm:"varchar(100) 'bank_name'" json:"bank_name" xml:"bank_name"`
	//
	CardNumber null.String `xorm:"varchar(100) 'card_number'" json:"card_number" xml:"card_number"`
	//
	RefType null.Int `xorm:"int(11) 'ref_type'" json:"ref_type" xml:"ref_type"`
	//
	NowCheckid null.String `xorm:"varchar(100) 'now_checkid'" json:"now_checkid" xml:"now_checkid"`
	//
	NowCheckname null.String `xorm:"varchar(100) 'now_checkname'" json:"now_checkname" xml:"now_checkname"`
	//
	NextCheckid null.String `xorm:"varchar(100) 'next_checkid'" json:"next_checkid" xml:"next_checkid"`
	//
	NextCheckname null.String `xorm:"varchar(100) 'next_checkname'" json:"next_checkname" xml:"next_checkname"`
}

// TableName table name of defined DepositOut
func (m *DepositOut) TableName() string {
	return "deposit_out"
}
