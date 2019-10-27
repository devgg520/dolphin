// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
)

// Tools struct
type Tools struct {
}

// Build func
func (tool *Tools) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
	}
	return []*gen.TmplCfg{
		&gen.TmplCfg{
			Text:     tempalte.TmplTools,
			FilePath: path.Join(dir, "utils", "tools"),
			Data:     data,
			Overlap:  gen.OverlapSkip,
		},
	}, nil
}