// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SatisfyVisitModelContent defined
type SatisfyVisitModelContent struct {
	//
	SVMCId null.Int `xorm:"int(11) pk notnull autoincr 's_v_m_c_id'" json:"s_v_m_c_id" xml:"s_v_m_c_id"`
	//
	SatisfyVisitCustomerType null.Int `xorm:"int(11) 'satisfy_visit_customer_type'" json:"satisfy_visit_customer_type" xml:"satisfy_visit_customer_type"`
	//
	SatisfyVisitType null.Int `xorm:"int(11) 'satisfy_visit_type'" json:"satisfy_visit_type" xml:"satisfy_visit_type"`
	//
	Content null.String `xorm:"varchar(500) 'content'" json:"content" xml:"content"`
	//
	SvmId null.Int `xorm:"int(11) 'svm_id'" json:"svm_id" xml:"svm_id"`
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
	Fraction null.Int `xorm:"int(11) 'fraction'" json:"fraction" xml:"fraction"`
}

// TableName table name of defined SatisfyVisitModelContent
func (m *SatisfyVisitModelContent) TableName() string {
	return "satisfy_visit_model_content"
}
