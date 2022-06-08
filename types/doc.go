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

//go:generate go run gen.go
