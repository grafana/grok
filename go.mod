module github.com/grafana/grok

go 1.19

require (
	cuelang.org/go v0.5.0-beta.2
	github.com/grafana/codejen v0.0.4-0.20221122220907-a5e7cc5407b3
	github.com/grafana/grafana v1.9.2-0.20230202155234-99a16d27c19b
	github.com/grafana/thema v0.0.0-20230208105941-fc71b67f3cca
)

require (
	github.com/cockroachdb/apd/v2 v2.0.2 // indirect
	github.com/dave/dst v0.27.2 // indirect
	github.com/deepmap/oapi-codegen v1.12.4 // indirect
	github.com/emicklei/proto v1.10.0 // indirect
	github.com/getkin/kin-openapi v0.112.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/grafana/cuetsy v0.1.5 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.10.0 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/protocolbuffers/txtpbfmt v0.0.0-20220428173112-74888fd59c2b // indirect
	github.com/rivo/uniseg v0.3.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	github.com/yalue/merged_fs v1.2.2 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.5.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.25.3 // indirect
	k8s.io/apimachinery v0.25.3 // indirect
	k8s.io/klog/v2 v2.80.0 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace github.com/grafana/grafana => ../grafana

replace github.com/grafana/codejen => ../codejen

replace github.com/grafana/thema => ../thema

replace cuelang.org/go => github.com/sdboyer/cue v0.5.0-beta.2.0.20221218111347-341999f48bdb

replace k8s.io/client-go => k8s.io/client-go v0.25.0
