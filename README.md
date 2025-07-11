# zerolint golangci Plugin

[![Go Reference](https://pkg.go.dev/badge/fillmore-labs.com/zerolint-golangci-plugin.svg)](https://pkg.go.dev/fillmore-labs.com/zerolint-golangci-plugin)
[![Test](https://github.com/fillmore-labs/zerolint-golangci-plugin/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/fillmore-labs/zerolint-golangci-plugin/actions/workflows/test.yml)
[![License](https://img.shields.io/github/license/fillmore-labs/zerolint-golangci-plugin)](https://www.apache.org/licenses/LICENSE-2.0)

## Usage

Add a file `.custom-gcl.yaml` to your source with

```YAML
---
version: v2.2.1
plugins:
  - module: fillmore-labs.com/zerolint-golangci-plugin
    version: v0.0.12
```

then run `golangci-lint custom`. You get an `custom-gcl` executable that can be configured in `.golangci.yaml`:

```YAML
---
version: "2"
linters:
  enable:
    - zerolint
  settings:
    custom:
      zerolint:
        type: module
        description: "zerolint checks usage patterns of pointers to zero-size types."
        original-url: "https://fillmore-labs.com/zerolint"
        settings:
          level: "full"
          excluded:
            - "struct{}"
          generated: true
          match: "^your\\.domain/package/path"
```

and can be used like `golangci-lint`:

```shell
./custom-gcl run .
```

See also the golangci-lint
[module plugin system documentation](https://golangci-lint.run/plugins/module-plugins/#the-automatic-way).
