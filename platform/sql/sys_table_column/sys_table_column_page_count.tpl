-- Code generated by dol build. Only Generate by tools if not existed.
select
    count(*) records
from
	sys_table_column
where
	sys_table_column.id {{.ne}} ""
	and
	sys_table_column.del_flag {{.ne}} 1
{{if ne .role_rule ""}}
	and {{.role_rule}}
{{end}}
