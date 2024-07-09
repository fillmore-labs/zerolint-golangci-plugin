package plugin

import (
	"fillmore-labs.com/zerolint/pkg/analyzer"
	"fillmore-labs.com/zerolint/pkg/visitor"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const Name = "zerolint"

func init() { //nolint:gochecknoinits
	register.Plugin(Name, New)
}

type Settings struct {
	Excluded []string `json:"excluded"`
}

func New(settings any) (register.LinterPlugin, error) { //nolint:ireturn
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return Plugin{settings: s}, nil
}

type Plugin struct {
	settings Settings
}

func (p Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	run := func(pass *analysis.Pass) (any, error) {
		excludes := make(map[string]struct{}, len(p.settings.Excluded))
		for _, ex := range p.settings.Excluded {
			excludes[ex] = struct{}{}
		}

		v := visitor.Visitor{Pass: pass, Excludes: excludes}
		v.Run()

		return any(nil), nil
	}

	analyzer := &analysis.Analyzer{
		Name:     Name,
		Doc:      analyzer.Doc,
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	return []*analysis.Analyzer{analyzer}, nil
}

func (Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
