package plugin

import (
	"fillmore-labs.com/zerolint/pkg/analyzer"
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
	Basic    bool     `json:"basic"`
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
	run := analyzer.NewRun(
		analyzer.WithExcludes(p.settings.Excluded),
		analyzer.WithBasic(p.settings.Basic),
		analyzer.WithGenerated(true),
	)

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
