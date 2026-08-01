// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "426d1bbc4004b17bc063d5e48caea6a025d707a74f693974032b90af20af7686"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c89e0a49deabdb19f2054369d4ccd61ee6667c03"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
