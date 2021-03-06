// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// NewStuTransition defined
type NewStuTransition struct {
	//
	NSTId null.Int `xorm:"int(11) pk notnull autoincr 'n_s_t_id'" json:"n_s_t_id" xml:"n_s_t_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	StuSchoolInfo null.String `xorm:"varchar(500) 'stu_school_info'" json:"stu_school_info" xml:"stu_school_info"`
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
	StuParentName null.String `xorm:"varchar(100) 'stu_parent_name'" json:"stu_parent_name" xml:"stu_parent_name"`
	//
	StuParentPhone null.Float `xorm:"float(50,2) 'stu_parent_phone'" json:"stu_parent_phone" xml:"stu_parent_phone"`
	//
	StuParentDetailInfo null.String `xorm:"varchar(500) 'stu_parent_detail_info'" json:"stu_parent_detail_info" xml:"stu_parent_detail_info"`
	//
	EcId null.Int `xorm:"int(11) 'ec_id'" json:"ec_id" xml:"ec_id"`
	//
	CpcId null.Int `xorm:"int(11) 'cpc_id'" json:"cpc_id" xml:"cpc_id"`
	//
	TextBookInfo null.String `xorm:"varchar(500) 'text_book_info'" json:"text_book_info" xml:"text_book_info"`
	//
	ClassTimeInfo null.String `xorm:"varchar(500) 'class_time_info'" json:"class_time_info" xml:"class_time_info"`
	//
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" xml:"tea_id"`
	//
	TaId null.Int `xorm:"int(11) 'ta_id'" json:"ta_id" xml:"ta_id"`
	//
	EcTransitionDate null.Time `xorm:"datetime 'ec_transition_date'" json:"ec_transition_date" xml:"ec_transition_date"`
	//
	StuCharacterHobby null.String `xorm:"varchar(1000) 'stu_character_hobby'" json:"stu_character_hobby" xml:"stu_character_hobby"`
	//
	StuLearnNeed null.String `xorm:"varchar(1000) 'stu_learn_need'" json:"stu_learn_need" xml:"stu_learn_need"`
	//
	StuDiscountInfo null.String `xorm:"varchar(1000) 'stu_discount_info'" json:"stu_discount_info" xml:"stu_discount_info"`
	//
	StuSpecialInfo null.String `xorm:"varchar(1000) 'stu_special_info'" json:"stu_special_info" xml:"stu_special_info"`
	//
	StuPromiseInfo null.String `xorm:"varchar(1000) 'stu_promise_info'" json:"stu_promise_info" xml:"stu_promise_info"`
	//
	StuTransitionType null.Int `xorm:"int(11) 'stu_transition_type'" json:"stu_transition_type" xml:"stu_transition_type"`
}

// TableName table name of defined NewStuTransition
func (m *NewStuTransition) TableName() string {
	return "new_stu_transition"
}
