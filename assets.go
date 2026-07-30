// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "84f63facc9694fe32bfd2f6f4113bb013cc948d350bcb61d72c961250e4da36d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a81b1066cda41bc7792e4b5417083fb8cda02eb1"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
