// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1f878a294f59e2f35233a8d53d23616a5a52bc2b63cc81d407cfda9bc5ca13a5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cfd6060c390ae735ae3f198dd0bb02a7f6720b72"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
