package template

// TmplSQLMap defined template
var TmplSQLMap = `
{{- $TableName := .Table.Name -}}
<sqlMap>
    <sql id="insert_{{$TableName}}">
        {{.Table.SQLInsertOne .Table $TableName}}
    </sql>
    <sql id="update_{{$TableName}}">
        {{.Table.SQLUpdateOne .Table $TableName}}
    </sql>
    <sql id="delete_{{$TableName}}">
        delete from {{$TableName}}
		where id =?id
    </sql>
    <sql id="selectone_{{$TableName}}">
        {{.Table.SQLSelectOne .Table $TableName}}
    </sql>
    <sql id="selectall_{{$TableName}}">
        {{.Table.SQLSelectAll .Table $TableName}}
    </sql>
</sqlMap>
`