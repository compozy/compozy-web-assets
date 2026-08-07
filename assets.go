// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0ef26f90e5d173fbd9e75c2b1f9e676497b911d15d7fc7e22f0fe183f6fe1e23"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f8117d9d04aec9501080f630575ef65c4e984344"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
