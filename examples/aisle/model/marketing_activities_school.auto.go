// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MarketingActivitiesSchool defined
type MarketingActivitiesSchool struct {
	//
	MASId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_s_id'" json:"m_a_s_id" xml:"m_a_s_id"`
	//
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" xml:"ma_id"`
	//
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" xml:"school_id"`
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

// TableName table name of defined MarketingActivitiesSchool
func (m *MarketingActivitiesSchool) TableName() string {
	return "marketing_activities_school"
}
