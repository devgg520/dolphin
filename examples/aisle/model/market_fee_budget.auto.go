// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketFeeBudget defined
type MarketFeeBudget struct {
	//
	MFBId null.Int `xorm:"int(11) pk notnull autoincr 'm_f_b_id'" json:"m_f_b_id" xml:"m_f_b_id"`
	//
	ProjectName null.Int `xorm:"int(11) 'project_name'" json:"project_name" xml:"project_name"`
	//
	Digest null.String `xorm:"varchar(2000) 'digest'" json:"digest" xml:"digest"`
	//
	BudgetDetail null.String `xorm:"varchar(2000) 'budget_detail'" json:"budget_detail" xml:"budget_detail"`
	//
	BgdgetMoney null.Float `xorm:"float(10,2) 'bgdget_money'" json:"bgdget_money" xml:"bgdget_money"`
	//
	Remake null.String `xorm:"varchar(2000) 'remake'" json:"remake" xml:"remake"`
	//
	Xgygout null.Float `xorm:"float(10,2) 'xgygout'" json:"xgygout" xml:"xgygout"`
	//
	Xgygin null.Float `xorm:"float(10,2) 'xgygin'" json:"xgygin" xml:"xgygin"`
	//
	Xgygbmr null.Float `xorm:"float(10,2) 'xgygbmr'" json:"xgygbmr" xml:"xgygbmr"`
	//
	BudgetDate null.Time `xorm:"datetime 'budget_date'" json:"budget_date" xml:"budget_date"`
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
	BudgetSchool null.Int `xorm:"int(11) 'budget_school'" json:"budget_school" xml:"budget_school"`
	//
	ProName null.Int `xorm:"int(11) 'pro_name'" json:"pro_name" xml:"pro_name"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	BeginMonth null.Time `xorm:"datetime 'begin_month'" json:"begin_month" xml:"begin_month"`
	//
	EndMonth null.Time `xorm:"datetime 'end_month'" json:"end_month" xml:"end_month"`
	//
	Xgygmoney null.Float `xorm:"float(10,2) 'xgygmoney'" json:"xgygmoney" xml:"xgygmoney"`
	//
	MfbCity null.Int `xorm:"int(11) 'mfb_city'" json:"mfb_city" xml:"mfb_city"`
}

// TableName table name of defined MarketFeeBudget
func (m *MarketFeeBudget) TableName() string {
	return "market_fee_budget"
}
