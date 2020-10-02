// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TrainerTable defined
type TrainerTable struct {
	//
	TTId null.Int `xorm:"int(11) pk notnull autoincr 't_t_id'" json:"t_t_id" xml:"t_t_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	TrainId null.Int `xorm:"int(11) 'train_id'" json:"train_id" xml:"train_id"`
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
	CppDesc null.String `xorm:"varchar(500) 'cpp_desc'" json:"cpp_desc" xml:"cpp_desc"`
	//
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" xml:"start_time"`
	//
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" xml:"end_time"`
	//
	TrainingState null.Int `xorm:"int(11) 'training_state'" json:"training_state" xml:"training_state"`
	//
	EvaluteRemark null.String `xorm:"varchar(300) 'evalute_remark'" json:"evalute_remark" xml:"evalute_remark"`
}

// TableName table name of defined TrainerTable
func (m *TrainerTable) TableName() string {
	return "trainer_table"
}