// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// Department defined
type Department struct {
	//
	DepartmentId null.Int `xorm:"int(11) pk notnull autoincr 'department_id'" json:"department_id" xml:"department_id"`
	//
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" xml:"parent_id"`
	//
	DepartmentName null.String `xorm:"varchar(100) 'department_name'" json:"department_name" xml:"department_name"`
	//
	DepartmentNumber null.String `xorm:"varchar(100) 'department_number'" json:"department_number" xml:"department_number"`
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
	SchoolId null.Int `xorm:"int(11) 'school_id'" json:"school_id" xml:"school_id"`
	//
	DepCity null.Int `xorm:"int(11) 'dep_city'" json:"dep_city" xml:"dep_city"`
}

// TableName table name of defined Department
func (m *Department) TableName() string {
	return "department"
}
