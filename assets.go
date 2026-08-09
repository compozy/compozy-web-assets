// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "949801abf9465d3a6ab087280c06cc82b3360c685197cb4c6977c432520b4855"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4aeb4cd3d63877fc28ef17a7533fce19102c072d"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
