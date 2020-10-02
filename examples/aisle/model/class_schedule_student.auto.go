// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// ClassScheduleStudent defined
type ClassScheduleStudent struct {
	//
	CSSId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_s_id'" json:"c_s_s_id" xml:"c_s_s_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" xml:"stu_id"`
	//
	StuStartDate null.Time `xorm:"datetime 'stu_start_date'" json:"stu_start_date" xml:"stu_start_date"`
	//
	StuEndDate null.Time `xorm:"datetime 'stu_end_date'" json:"stu_end_date" xml:"stu_end_date"`
	//
	KqState null.Int `xorm:"int(11) 'kq_state'" json:"kq_state" xml:"kq_state"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" xml:"cs_id"`
	//
	KqType null.Int `xorm:"int(11) 'kq_type'" json:"kq_type" xml:"kq_type"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
	//
	ScsId null.Int `xorm:"int(11) 'scs_id'" json:"scs_id" xml:"scs_id"`
	//
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" xml:"if_kou_hour"`
	//
	NotKouReason null.String `xorm:"varchar(1000) 'not_kou_reason'" json:"not_kou_reason" xml:"not_kou_reason"`
	//
	KqKc null.Float `xorm:"float(10,2) 'kq_kc'" json:"kq_kc" xml:"kq_kc"`
	//
	KqHour null.Float `xorm:"float(10,2) 'kq_hour'" json:"kq_hour" xml:"kq_hour"`
	//
	ClassDate null.Time `xorm:"datetime 'class_date'" json:"class_date" xml:"class_date"`
	//
	ClassBeginTime null.Time `xorm:"datetime 'class_begin_time'" json:"class_begin_time" xml:"class_begin_time"`
	//
	ClassEndTime null.Time `xorm:"datetime 'class_end_time'" json:"class_end_time" xml:"class_end_time"`
	//
	KfId null.Int `xorm:"int(11) 'kf_id'" json:"kf_id" xml:"kf_id"`
	//
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" xml:"class_id"`
	//
	KxPrice null.Float `xorm:"float(10,2) 'kx_price'" json:"kx_price" xml:"kx_price"`
	//
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" xml:"check_state"`
	//
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" xml:"check_user"`
	//
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" xml:"check_time"`
	//
	CheckRemark null.String `xorm:"varchar(2000) 'check_remark'" json:"check_remark" xml:"check_remark"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" xml:"buss_type"`
	//
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" xml:"sct_id"`
	//
	SingleAllKxmoney null.Float `xorm:"float(50,2) 'single_all_kxmoney'" json:"single_all_kxmoney" xml:"single_all_kxmoney"`
	//
	SbtId null.Int `xorm:"int(11) 'sbt_id'" json:"sbt_id" xml:"sbt_id"`
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" xml:"sch_id"`
	//
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" xml:"par_id"`
	//
	IfUpdateKf null.Int `xorm:"int(11) 'if_update_kf'" json:"if_update_kf" xml:"if_update_kf"`
	//
	JfPrice null.Float `xorm:"float(50,2) 'jf_price'" json:"jf_price" xml:"jf_price"`
	//
	JfAllprice null.Float `xorm:"float(50,2) 'jf_allprice'" json:"jf_allprice" xml:"jf_allprice"`
	//
	IfMz null.Int `xorm:"int(11) 'if_mz'" json:"if_mz" xml:"if_mz"`
	//
	DetailHour null.Float `xorm:"float(50,2) 'detail_hour'" json:"detail_hour" xml:"detail_hour"`
	//
	UseMz null.Float `xorm:"float(50,2) 'use_mz'" json:"use_mz" xml:"use_mz"`
	//
	MzReason null.String `xorm:"varchar(20) 'mz_reason'" json:"mz_reason" xml:"mz_reason"`
	//
	ZdReason null.Int `xorm:"int(11) 'zd_reason'" json:"zd_reason" xml:"zd_reason"`
	//
	ZdRsType null.Int `xorm:"int(11) 'zd_rs_type'" json:"zd_rs_type" xml:"zd_rs_type"`
	//
	PkOf null.Int `xorm:"int(11) 'pk_of'" json:"pk_of" xml:"pk_of"`
	//
	GdkbIfUse null.Int `xorm:"int(11) 'gdkb_if_use'" json:"gdkb_if_use" xml:"gdkb_if_use"`
	//
	GdkbUseDate null.Time `xorm:"datetime 'gdkb_use_date'" json:"gdkb_use_date" xml:"gdkb_use_date"`
	//
	TurnInSch null.Int `xorm:"int(11) 'turn_in_sch'" json:"turn_in_sch" xml:"turn_in_sch"`
	//
	NextCssId null.Int `xorm:"int(11) 'next_css_id'" json:"next_css_id" xml:"next_css_id"`
}

// TableName table name of defined ClassScheduleStudent
func (m *ClassScheduleStudent) TableName() string {
	return "class_schedule_student"
}