// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1f18821fbc30d22159e4667370cf840a7c34d75de28c28ec697dbaa1b8ff5edf"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e656bed8bba0cc0b159772a9c9db2179afd71fb2"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
