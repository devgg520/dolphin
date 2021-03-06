package pipe

import "github.com/2637309949/dolphin/cmd/dolphin/schema"

type (
	// Overlap int
	Overlap int
	// TmplCfg struct
	TmplCfg struct {
		Text     string
		FilePath string
		Data     interface{}
		Overlap  Overlap
		GOFmt    bool
		GOProto  bool
		Format   func(string) string
	}
	// Pipe interface
	Pipe interface {
		Name() string
		Build(string, []string, *schema.Application) ([]*TmplCfg, error)
	}
)

// OverlapSkip Overlap
const (
	OverlapSkip Overlap = iota + 1
	OverlapWrite
	OverlapInc
)
