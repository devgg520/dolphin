// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// AutoCsTimeSet defined
type AutoCsTimeSet struct {
	//
	ACTSId null.Int `xorm:"int(11) pk notnull autoincr 'a_c_t_s_id'" json:"a_c_t_s_id" xml:"a_c_t_s_id"`
	//
	SetWeek null.Int `xorm:"int(11) 'set_week'" json:"set_week" xml:"set_week"`
	//
	SetHour null.Float `xorm:"float(11,2) 'set_hour'" json:"set_hour" xml:"set_hour"`
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
	SetMinute null.Int `xorm:"int(11) 'set_minute'" json:"set_minute" xml:"set_minute"`
	//
	IfOpen null.Int `xorm:"int(11) 'if_open'" json:"if_open" xml:"if_open"`
	//
	TimeName null.String `xorm:"varchar(20) 'time_name'" json:"time_name" xml:"time_name"`
	//
	StartTimeStr null.String `xorm:"varchar(5) 'start_time_str'" json:"start_time_str" xml:"start_time_str"`
	//
	EndTimeStr null.String `xorm:"varchar(5) 'end_time_str'" json:"end_time_str" xml:"end_time_str"`
}

// TableName table name of defined AutoCsTimeSet
func (m *AutoCsTimeSet) TableName() string {
	return "auto_cs_time_set"
}
