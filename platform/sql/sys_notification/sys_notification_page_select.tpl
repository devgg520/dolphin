select
    sys_notification.id
from
	sys_notification
where
	sys_notification.id {{.ne}} ""
	and
	sys_notification.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
LIMIT {{.size}} OFFSET {{.offset}}
