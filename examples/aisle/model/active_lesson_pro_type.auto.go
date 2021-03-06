// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ActiveLessonProType defined
type ActiveLessonProType struct {
	//
	ALPTId null.Int `xorm:"int(11) pk notnull autoincr 'a_l_p_t_id'" json:"a_l_p_t_id" xml:"a_l_p_t_id"`
	//
	ActiveId null.Int `xorm:"int(11) 'active_id'" json:"active_id" xml:"active_id"`
	//
	ProTypeId null.Int `xorm:"int(11) 'pro_type_id'" json:"pro_type_id" xml:"pro_type_id"`
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

// TableName table name of defined ActiveLessonProType
func (m *ActiveLessonProType) TableName() string {
	return "active_lesson_pro_type"
}
