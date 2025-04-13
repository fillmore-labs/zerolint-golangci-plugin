// Copyright 2024 Oliver Eikemeier. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fillmore-labs.com/zerolint/pkg/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// Name is the linters name.
const Name = "zerolint"

func init() { //nolint:gochecknoinits
	register.Plugin(Name, New)
}

// Settings are the linters settings.
type Settings struct {
	Excluded []string `json:"excluded"`
	Full     bool     `json:"full"`
}

// New creates a new [Plugin] instance with the given [Settings].
func New(settings any) (register.LinterPlugin, error) { //nolint:ireturn
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return Plugin{settings: s}, nil
}

// Plugin is a zerolint linter as a [register.LinterPlugin].
type Plugin struct {
	settings Settings
}

// BuildAnalyzers returns the [analysis.Analyzer]s for a zerolint run.
func (p Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	run := analyzer.NewRun(
		analyzer.WithExcludes(p.settings.Excluded),
		analyzer.WithFull(p.settings.Full),
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

// GetLoadMode returns the golangci load mode.
func (Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
