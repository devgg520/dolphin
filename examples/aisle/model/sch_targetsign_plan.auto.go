// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SchTargetsignPlan defined
type SchTargetsignPlan struct {
	//
	STPId null.Int `xorm:"int(11) pk notnull autoincr 's_t_p_id'" json:"s_t_p_id" xml:"s_t_p_id"`
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" xml:"sch_id"`
	//
	MarkInNum null.Float `xorm:"float(11,2) 'mark_in_num'" json:"mark_in_num" xml:"mark_in_num"`
	//
	MarkOutNum null.Float `xorm:"float(50,2) 'mark_out_num'" json:"mark_out_num" xml:"mark_out_num"`
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
	NetworkTargetsignNum null.Float `xorm:"float(50,2) 'network_targetsign_num'" json:"network_targetsign_num" xml:"network_targetsign_num"`
	//
	LtTargetsignNum null.Float `xorm:"float(50,2) 'lt_targetsign_num'" json:"lt_targetsign_num" xml:"lt_targetsign_num"`
	//
	PpTargetsignNum null.Float `xorm:"float(50,2) 'pp_targetsign_num'" json:"pp_targetsign_num" xml:"pp_targetsign_num"`
	//
	QdTargetsignNum null.Float `xorm:"float(50,2) 'qd_targetsign_num'" json:"qd_targetsign_num" xml:"qd_targetsign_num"`
	//
	QtTargetsignNum null.Float `xorm:"float(50,2) 'qt_targetsign_num'" json:"qt_targetsign_num" xml:"qt_targetsign_num"`
	//
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" xml:"plan_month"`
	//
	AllQdNum null.Float `xorm:"float(50,2) 'all_qd_num'" json:"all_qd_num" xml:"all_qd_num"`
	//
	AllQdMoney null.Float `xorm:"float(50,2) 'all_qd_money'" json:"all_qd_money" xml:"all_qd_money"`
	//
	AllTzMoney null.Float `xorm:"float(50,2) 'all_tz_money'" json:"all_tz_money" xml:"all_tz_money"`
}

// TableName table name of defined SchTargetsignPlan
func (m *SchTargetsignPlan) TableName() string {
	return "sch_targetsign_plan"
}
