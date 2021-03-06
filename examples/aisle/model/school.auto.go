// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// School defined
type School struct {
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	SchoolId null.Int `xorm:"int(11) pk notnull autoincr 'school_id'" json:"school_id" xml:"school_id"`
	//
	SchoolName null.String `xorm:"varchar(500) 'school_name'" json:"school_name" xml:"school_name"`
	//
	SchoolArea null.String `xorm:"varchar(2000) 'school_area'" json:"school_area" xml:"school_area"`
	//
	SchoolRemark null.String `xorm:"varchar(2000) 'school_remark'" json:"school_remark" xml:"school_remark"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" xml:"organ_id"`
}

// TableName table name of defined School
func (m *School) TableName() string {
	return "school"
}
