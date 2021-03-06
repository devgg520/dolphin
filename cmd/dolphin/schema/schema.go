// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package schema

// Success struct
type Success struct {
	Common
	Type string
}

// Failure struct
type Failure struct {
	Common
	Type string
}

// Return struct
type Return struct {
	Common
	Success *Success
	Failure *Failure
}

// Param struct
type Param struct {
	Common
	Type  string
	Value string
}

// API struct
type API struct {
	Common
	Roles   []string
	Auth    bool
	Version string
	Path    string
	Func    string
	Table   string
	Method  string
	Cache   uint64
	Params  []*Param
	Return  *Return
}

// Controller deifned
type Controller struct {
	Common
	APIS   []*API
	Prefix string
}

// Service defined
type Service struct {
	Common
	RPCS []*RPC
}

// RPC defined
type RPC struct {
	Common
	Request *Request
	Reply   *Reply
}

// Request defined
type Request struct {
	Common
	Type string
}

// Reply defined
type Reply struct {
	Common
	Type string
}

// Prop struct
type Prop struct {
	Common
	Type    string
	JSON    string
	Form    string
	Example string
}

// Bean struct
type Bean struct {
	Common
	Packages string
	Props    []*Prop
	Extends  string
}

// Column struct
type Column struct {
	Common
	Type    string
	Xorm    string
	JSON    string
	Form    string
	Example string
}

// Table struct
type Table struct {
	Common
	Bind     string
	Packages string
	Columns  []*Column
	Extends  string
}

// Application struct
type Application struct {
	Common
	PackageName string `validate:"required"`
	Controllers []*Controller
	Services    []*Service
	Beans       []*Bean
	Tables      []*Table
}
