// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// AddAgreementMb defined
type AddAgreementMb struct {
	//
	AAMId null.Int `xorm:"int(11) pk notnull autoincr 'a_a_m_id'" json:"a_a_m_id" xml:"a_a_m_id"`
	//
	AamName null.String `xorm:"varchar(20) 'aam_name'" json:"aam_name" xml:"aam_name"`
	//
	AamHead null.Int `xorm:"int(11) 'aam_head'" json:"aam_head" xml:"aam_head"`
	//
	AamMidd null.Int `xorm:"int(11) 'aam_midd'" json:"aam_midd" xml:"aam_midd"`
	//
	AamLast null.Int `xorm:"int(11) 'aam_last'" json:"aam_last" xml:"aam_last"`
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
	AamTitle null.String `xorm:"varchar(500) 'aam_title'" json:"aam_title" xml:"aam_title"`
	//
	OpenOrClose null.Int `xorm:"int(11) 'open_or_close'" json:"open_or_close" xml:"open_or_close"`
}

// TableName table name of defined AddAgreementMb
func (m *AddAgreementMb) TableName() string {
	return "add_agreement_mb"
}
