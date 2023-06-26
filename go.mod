module github.com/grafana/grok

go 1.19

require (
	cuelang.org/go v0.5.0
	github.com/grafana/codejen v0.0.4-0.20221122220907-a5e7cc5407b3
	github.com/grafana/grafana v1.9.2-0.20230609164553-81d85ce8028d
	github.com/grafana/kindsys v0.0.0-20230427152021-bb328815be7a
	github.com/grafana/thema v0.0.0-20230417103609-99b482c479fe
)

require (
	github.com/cockroachdb/apd/v2 v2.0.2 // indirect
	github.com/dave/dst v0.27.2 // indirect
	github.com/deepmap/oapi-codegen v1.12.4 // indirect
	github.com/emicklei/proto v1.10.0 // indirect
	github.com/getkin/kin-openapi v0.115.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/grafana/cuetsy v0.1.9 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.10.2 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/perimeterx/marshmallow v1.1.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/protocolbuffers/txtpbfmt v0.0.0-20220428173112-74888fd59c2b // indirect
	github.com/rivo/uniseg v0.3.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	github.com/yalue/merged_fs v1.2.2 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/grafana/grafana => ../grafana

// replace github.com/grafana/codejen => ../codejen

// replace github.com/grafana/thema => ../thema

replace cuelang.org/go => github.com/sdboyer/cue v0.5.0-beta.2.0.20221218111347-341999f48bdb

replace github.com/deepmap/oapi-codegen => github.com/spinillos/oapi-codegen v1.12.5-0.20230417081915-2945b61c0b1c
