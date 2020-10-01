// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuCtTimeRange defined
type StuCtTimeRange struct {
	//
	SCTRId null.Int `xorm:"int(11) pk notnull autoincr 's_c_t_r_id'" json:"s_c_t_r_id" xml:"s_c_t_r_id"`
	//
	TrId null.Int `xorm:"int(11) 'tr_id'" json:"tr_id" xml:"tr_id"`
	//
	StartDate null.Time `xorm:"datetime 'start_date'" json:"start_date" xml:"start_date"`
	//
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" xml:"end_date"`
	//
	StartWeek null.Int `xorm:"int(11) 'start_week'" json:"start_week" xml:"start_week"`
	//
	EndWeek null.Int `xorm:"int(11) 'end_week'" json:"end_week" xml:"end_week"`
	//
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" xml:"start_time"`
	//
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" xml:"end_time"`
	//
	Hour null.Float `xorm:"float(10,2) 'hour'" json:"hour" xml:"hour"`
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
	TimeRemark null.String `xorm:"varchar(500) 'time_remark'" json:"time_remark" xml:"time_remark"`
}

// TableName table name of defined StuCtTimeRange
func (m *StuCtTimeRange) TableName() string {
	return "stu_ct_time_range"
}
