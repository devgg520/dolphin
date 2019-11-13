// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tempalte

// TmplMain defined template
var TmplMain = `// Code generated by dol build. Only Generate by tools if not existed.
// source: main.go

package main

import (
	_ "{{.PackageName}}/app"
	_ "{{.PackageName}}/model"

	"github.com/2637309949/dolphin/srv/cli"
	_ "github.com/go-sql-driver/mysql"
)

// @title {{.Name}} api doc
// @version 1.0
// @description This is a {{.Name}} server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name api support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8089
// @BasePath /
func main() {
	cli.Run()
}
`
