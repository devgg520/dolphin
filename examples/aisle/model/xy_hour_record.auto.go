// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
)

// XyHourRecord defined
type XyHourRecord struct {
	//
	XHRId null.Int `xorm:"int(11) pk notnull autoincr 'x_h_r_id'" json:"x_h_r_id" xml:"x_h_r_id"`
	//
	PkStu null.Int `xorm:"int(11) 'pk_stu'" json:"pk_stu" xml:"pk_stu"`
	//
	CtName null.String `xorm:"varchar(500) 'ct_name'" json:"ct_name" xml:"ct_name"`
	//
	SurplusHour null.Float `xorm:"float(11,2) 'surplus_hour'" json:"surplus_hour" xml:"surplus_hour"`
	//
	FinalMoney decimal.Decimal `xorm:"decimal(11,2) 'final_money'" json:"final_money" xml:"final_money"`
	//
	OfStartDate null.Time `xorm:"datetime 'of_start_date'" json:"of_start_date" xml:"of_start_date"`
	//
	OfEndDate null.Time `xorm:"datetime 'of_end_date'" json:"of_end_date" xml:"of_end_date"`
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
	OfBatch null.Int `xorm:"int(11) 'of_batch'" json:"of_batch" xml:"of_batch"`
	//
	CtCourseType null.Int `xorm:"int(11) 'ct_course_type'" json:"ct_course_type" xml:"ct_course_type"`
	//
	OfCreateDate null.Time `xorm:"datetime 'of_create_date'" json:"of_create_date" xml:"of_create_date"`
	//
	PkOf null.Int `xorm:"int(11) 'pk_of'" json:"pk_of" xml:"pk_of"`
	//
	SctIds null.String `xorm:"varchar(500) 'sct_ids'" json:"sct_ids" xml:"sct_ids"`
}

// TableName table name of defined XyHourRecord
func (m *XyHourRecord) TableName() string {
	return "xy_hour_record"
}