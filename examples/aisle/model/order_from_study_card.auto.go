// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// OrderFromStudyCard defined
type OrderFromStudyCard struct {
	//
	OFSCId null.Int `xorm:"int(11) pk notnull autoincr 'o_f_s_c_id'" json:"o_f_s_c_id" xml:"o_f_s_c_id"`
	//
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" xml:"of_id"`
	//
	ScId null.Int `xorm:"int(11) 'sc_id'" json:"sc_id" xml:"sc_id"`
	//
	DiscountMoney null.Float `xorm:"float(11,2) 'discount_money'" json:"discount_money" xml:"discount_money"`
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
}

// TableName table name of defined OrderFromStudyCard
func (m *OrderFromStudyCard) TableName() string {
	return "order_from_study_card"
}
