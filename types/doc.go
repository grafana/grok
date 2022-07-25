/* Package static is the parent directory to all the Go types generated from the
canonical Grafana schemas.

Each subdirectory contains the types for one Grafana entity type. These types
are organized into subpackages, with each subpackage containing the exported types
for one major version (aka Thema sequence).

These pure Go types are simpler and more familiar to use for most Go developers
than the hybrid Thema types. They are also clunkier, offer weaker validation, and are
less expressive.
*/

package static

import (
	_ "cuelang.org/go/cue/cuecontext"
	_ "cuelang.org/go/pkg/encoding/yaml"
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
	_ "github.com/getkin/kin-openapi/openapi3"
	_ "github.com/grafana/grafana/pkg/codegen"
	_ "github.com/grafana/grafana/pkg/framework/coremodel"
	_ "github.com/grafana/grafana/pkg/framework/coremodel/registry"
	_ "github.com/grafana/thema"
	_ "github.com/grafana/thema/encoding/openapi"
	_ "golang.org/x/tools/imports"
)

//go:generate go run gen.go
