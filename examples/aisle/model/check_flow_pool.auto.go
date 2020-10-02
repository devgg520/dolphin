// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// CheckFlowPool defined
type CheckFlowPool struct {
	//
	CFPId null.Int `xorm:"int(11) pk notnull autoincr 'c_f_p_id'" json:"c_f_p_id" xml:"c_f_p_id"`
	//
	FlowName null.String `xorm:"varchar(500) 'flow_name'" json:"flow_name" xml:"flow_name"`
	//
	ZdCheckState null.Int `xorm:"int(11) 'zd_check_state'" json:"zd_check_state" xml:"zd_check_state"`
	//
	PkCheckUser null.Int `xorm:"int(11) 'pk_check_user'" json:"pk_check_user" xml:"pk_check_user"`
	//
	CheckDate null.Time `xorm:"datetime 'check_date'" json:"check_date" xml:"check_date"`
	//
	NowCheckUser null.String `xorm:"varchar(500) 'now_check_user'" json:"now_check_user" xml:"now_check_user"`
	//
	NowFloor null.Int `xorm:"int(11) 'now_floor'" json:"now_floor" xml:"now_floor"`
	//
	TurnFloor null.Int `xorm:"int(11) 'turn_floor'" json:"turn_floor" xml:"turn_floor"`
	//
	HistoryCheckUser null.String `xorm:"varchar(500) 'history_check_user'" json:"history_check_user" xml:"history_check_user"`
	//
	PkFlowSet null.Int `xorm:"int(11) 'pk_flow_set'" json:"pk_flow_set" xml:"pk_flow_set"`
	//
	PkRefundId null.Int `xorm:"int(11) 'pk_refund_id'" json:"pk_refund_id" xml:"pk_refund_id"`
	//
	PkDOId null.Int `xorm:"int(11) 'pk_d_o_id'" json:"pk_d_o_id" xml:"pk_d_o_id"`
	//
	PkTSCId null.Int `xorm:"int(11) 'pk_t_s_c_id'" json:"pk_t_s_c_id" xml:"pk_t_s_c_id"`
}

// TableName table name of defined CheckFlowPool
func (m *CheckFlowPool) TableName() string {
	return "check_flow_pool"
}