// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StudentTimeResource defined
type StudentTimeResource struct {
	//
	STRId null.Int `xorm:"int(11) pk notnull autoincr 's_t_r_id'" json:"s_t_r_id" xml:"s_t_r_id"`
	//
	DataId null.Float `xorm:"float(11,2) 'data_id'" json:"data_id" xml:"data_id"`
	//
	TableId null.Float `xorm:"float(11,2) 'table_id'" json:"table_id" xml:"table_id"`
	//
	MainDataId null.Int `xorm:"int(11) 'main_data_id'" json:"main_data_id" xml:"main_data_id"`
	//
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" xml:"start_time"`
	//
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" xml:"end_time"`
	//
	ConflictContent null.String `xorm:"varchar(500) 'conflict_content'" json:"conflict_content" xml:"conflict_content"`
	//
	ShowContent null.String `xorm:"varchar(500) 'show_content'" json:"show_content" xml:"show_content"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	CourseType null.Int `xorm:"int(11) 'course_type'" json:"course_type" xml:"course_type"`
	//
	CourseState null.Int `xorm:"int(11) 'course_state'" json:"course_state" xml:"course_state"`
}

// TableName table name of defined StudentTimeResource
func (m *StudentTimeResource) TableName() string {
	return "student_time_resource"
}
