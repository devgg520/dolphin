// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// SystemOperaLogs defined
type SystemOperaLogs struct {
	//
	T120 null.Int `xorm:"int(11) pk notnull autoincr 't_12_0'" json:"t_12_0" xml:"t_12_0"`
	//
	DataId null.Float `xorm:"float(11,2) 'data_id'" json:"data_id" xml:"data_id"`
	//
	DataName null.String `xorm:"varchar(2000) 'data_name'" json:"data_name" xml:"data_name"`
	//
	UserIp null.String `xorm:"varchar(2000) 'user_ip'" json:"user_ip" xml:"user_ip"`
	//
	OperaObject null.String `xorm:"varchar(2000) 'opera_object'" json:"opera_object" xml:"opera_object"`
	//
	OperaType null.Int `xorm:"int(11) 'opera_type'" json:"opera_type" xml:"opera_type"`
	//
	ProgramIp null.String `xorm:"varchar(500) 'program_ip'" json:"program_ip" xml:"program_ip"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" xml:"update_date"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" xml:"updater"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined SystemOperaLogs
func (m *SystemOperaLogs) TableName() string {
	return "system_opera_logs"
}
