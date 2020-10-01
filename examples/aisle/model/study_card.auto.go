// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudyCard defined
type StudyCard struct {
	//
	StudyCardId null.Int `xorm:"int(11) pk notnull autoincr 'study_card_id'" json:"study_card_id" xml:"study_card_id"`
	//
	StudyCardnumber null.String `xorm:"varchar(100) 'study_cardnumber'" json:"study_cardnumber" xml:"study_cardnumber"`
	//
	ArrivedMoney null.Float `xorm:"float(50,2) 'arrived_money'" json:"arrived_money" xml:"arrived_money"`
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
	ArrivedType null.Int `xorm:"int(11) 'arrived_type'" json:"arrived_type" xml:"arrived_type"`
	//
	ArrivedDiscount null.Float `xorm:"float(50,2) 'arrived_discount'" json:"arrived_discount" xml:"arrived_discount"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" xml:"buss_type"`
	//
	CardName null.String `xorm:"varchar(300) 'card_name'" json:"card_name" xml:"card_name"`
}

// TableName table name of defined StudyCard
func (m *StudyCard) TableName() string {
	return "study_card"
}
