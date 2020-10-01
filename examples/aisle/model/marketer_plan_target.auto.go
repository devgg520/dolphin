// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketerPlanTarget defined
type MarketerPlanTarget struct {
	//
	MPTId null.Int `xorm:"int(11) pk notnull autoincr 'm_p_t_id'" json:"m_p_t_id" xml:"m_p_t_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	InPlan null.Float `xorm:"float(11,2) 'in_plan'" json:"in_plan" xml:"in_plan"`
	//
	OutPlan null.Float `xorm:"float(11,2) 'out_plan'" json:"out_plan" xml:"out_plan"`
	//
	DemoPlan null.Float `xorm:"float(11,2) 'demo_plan'" json:"demo_plan" xml:"demo_plan"`
	//
	AchievementPlan null.Float `xorm:"float(11,2) 'achievement_plan'" json:"achievement_plan" xml:"achievement_plan"`
	//
	PlanMonth null.Time `xorm:"datetime notnull default('1970-01-01 00:00:00') 'plan_month'" json:"plan_month" xml:"plan_month"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime notnull default('1970-01-01 00:00:00') 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime notnull default('1970-01-01 00:00:00') 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" xml:"pk_sch"`
}

// TableName table name of defined MarketerPlanTarget
func (m *MarketerPlanTarget) TableName() string {
	return "marketer_plan_target"
}
