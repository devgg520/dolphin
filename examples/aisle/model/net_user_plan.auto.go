// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
)

// NetUserPlan defined
type NetUserPlan struct {
	//
	NUPId null.Int `xorm:"int(11) pk notnull autoincr 'n_u_p_id'" json:"n_u_p_id" xml:"n_u_p_id"`
	//
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" xml:"pk_sch"`
	//
	PkPlaner null.Int `xorm:"int(11) 'pk_planer'" json:"pk_planer" xml:"pk_planer"`
	//
	InPlan null.Int `xorm:"int(11) 'in_plan'" json:"in_plan" xml:"in_plan"`
	//
	OutPlan null.Int `xorm:"int(11) 'out_plan'" json:"out_plan" xml:"out_plan"`
	//
	DemoPlan null.Int `xorm:"int(11) 'demo_plan'" json:"demo_plan" xml:"demo_plan"`
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
	AchievementPlan decimal.Decimal `xorm:"decimal(50,3) 'achievement_plan'" json:"achievement_plan" xml:"achievement_plan"`
	//
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" xml:"plan_month"`
}

// TableName table name of defined NetUserPlan
func (m *NetUserPlan) TableName() string {
	return "net_user_plan"
}
