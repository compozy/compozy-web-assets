// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c732906ca4291918c7f13c7059d608a4e4e1e66747696bc439df433e9588c6a0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d30b810ac5686c9d7bff8051eb0f424e80db5d5e"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
