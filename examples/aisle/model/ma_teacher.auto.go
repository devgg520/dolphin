// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// MaTeacher defined
type MaTeacher struct {
	//
	MaTeacherId null.Int `xorm:"int(11) pk notnull autoincr 'ma_teacher_id'" json:"ma_teacher_id" xml:"ma_teacher_id"`
	//
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" xml:"ma_id"`
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
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" xml:"user_id"`
	//
	TeaClassDate null.Time `xorm:"datetime 'tea_class_date'" json:"tea_class_date" xml:"tea_class_date"`
	//
	TeaBegintime null.Time `xorm:"datetime 'tea_begintime'" json:"tea_begintime" xml:"tea_begintime"`
	//
	TeaEndtime null.Time `xorm:"datetime 'tea_endtime'" json:"tea_endtime" xml:"tea_endtime"`
	//
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" xml:"if_kou_hour"`
	//
	ClassHour null.Float `xorm:"float(50,2) 'class_hour'" json:"class_hour" xml:"class_hour"`
}

// TableName table name of defined MaTeacher
func (m *MaTeacher) TableName() string {
	return "ma_teacher"
}
