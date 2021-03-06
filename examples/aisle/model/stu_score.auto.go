// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// StuScore defined
type StuScore struct {
	//
	StuScoreId null.Int `xorm:"int(11) pk notnull autoincr 'stu_score_id'" json:"stu_score_id" xml:"stu_score_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" xml:"of_id"`
	//
	TmId null.Int `xorm:"int(11) 'tm_id'" json:"tm_id" xml:"tm_id"`
	//
	LevelId null.Int `xorm:"int(11) 'level_id'" json:"level_id" xml:"level_id"`
	//
	UnitId null.Int `xorm:"int(11) 'unit_id'" json:"unit_id" xml:"unit_id"`
	//
	TestTime null.Time `xorm:"datetime 'test_time'" json:"test_time" xml:"test_time"`
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
	AllScore null.Float `xorm:"float(50,2) 'all_score'" json:"all_score" xml:"all_score"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	TauserId null.Int `xorm:"int(11) 'tauser_id'" json:"tauser_id" xml:"tauser_id"`
}

// TableName table name of defined StuScore
func (m *StuScore) TableName() string {
	return "stu_score"
}
