// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9dacebfca0eb9e320e187d2d2b83ea4eff8750309ff856d89b712f2fb4a86d0e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "80f51609ec25e11f3f8c380df35c37573f0d70fb"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
