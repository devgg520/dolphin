// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/gen/template"
	"github.com/2637309949/dolphin/client/schema"
	"github.com/2637309949/dolphin/packages/viper"
)

// SQLMap struct
type SQLMap struct {
}

// Name defined pipe name
func (app *SQLMap) Name() string {
	return "sqlmap"
}

// Build func
func (app *SQLMap) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, t := range node.Tables {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Application": node,
			"Table":       t,
		}
		tmplCfg := &gen.TmplCfg{
			Text:     template.TmplSQLMap,
			FilePath: path.Join(dir, viper.GetString("dir.sql"), viper.GetString("dir.sqlmap"), t.Name),
			Data:     data,
			Overlap:  gen.OverlapWrite,
			Suffix:   ".xml",
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
