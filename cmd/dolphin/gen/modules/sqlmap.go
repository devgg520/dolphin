// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// SQLMap struct
type SQLMap struct {
}

// Name defined pipe name
func (app *SQLMap) Name() string {
	return "sqlmap"
}

// Build func
func (app *SQLMap) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	sqlmapByte, _ := vfsutil.ReadFile(template.Assets, "sqlmap.tmpl")
	for _, t := range node.Tables {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Application": node,
			"Table":       t,
			"Viper":       viper.GetViper(),
		}
		filename := utils.FileNameTrimSuffix(t.Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(sqlmapByte),
			FilePath: path.Join(dir, viper.GetString("dir.sql"), viper.GetString("dir.sqlmap"), filename+".xml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
