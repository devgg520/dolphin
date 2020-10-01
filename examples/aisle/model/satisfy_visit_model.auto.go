// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SatisfyVisitModel defined
type SatisfyVisitModel struct {
	//
	SVMId null.Int `xorm:"int(11) pk notnull autoincr 's_v_m_id'" json:"s_v_m_id" xml:"s_v_m_id"`
	//
	ModelName null.String `xorm:"varchar(500) 'model_name'" json:"model_name" xml:"model_name"`
	//
	ModelDesc null.String `xorm:"varchar(500) 'model_desc'" json:"model_desc" xml:"model_desc"`
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

// TableName table name of defined SatisfyVisitModel
func (m *SatisfyVisitModel) TableName() string {
	return "satisfy_visit_model"
}
