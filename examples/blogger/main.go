// Code generated by dol build. Only Generate by tools if not existed.
// source: main.go

package main

import (
	// "github.com/2637309949/dolphin/platform/app" init
	_ "github.com/2637309949/dolphin/platform/app"
	// "github.com/go-sql-driver/mysql" init
	_ "github.com/go-sql-driver/mysql"
	// "blogger/app" init
	"blogger/app"
)

//go:generate dolphin build
func main() {
	app.Run()
}
