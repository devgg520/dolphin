// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SaleLanguageModel defined
type SaleLanguageModel struct {
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	BusinessType null.Int `xorm:"int(11) 'business_type'" json:"business_type" xml:"business_type"`
	//
	SLMId null.Int `xorm:"int(11) pk notnull autoincr 's_l_m_id'" json:"s_l_m_id" xml:"s_l_m_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" xml:"pt_id"`
	//
	SlmName null.String `xorm:"varchar(50) 'slm_name'" json:"slm_name" xml:"slm_name"`
	//
	SlmContent null.String `xorm:"varchar(500) 'slm_content'" json:"slm_content" xml:"slm_content"`
	//
	VisitStage null.Int `xorm:"int(11) 'visit_stage'" json:"visit_stage" xml:"visit_stage"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	SlmType null.Int `xorm:"int(11) 'slm_type'" json:"slm_type" xml:"slm_type"`
	//
	StuSystemSta null.Int `xorm:"int(11) 'stu_system_sta'" json:"stu_system_sta" xml:"stu_system_sta"`
}

// TableName table name of defined SaleLanguageModel
func (m *SaleLanguageModel) TableName() string {
	return "sale_language_model"
}
