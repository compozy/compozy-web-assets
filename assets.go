// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "426d1bbc4004b17bc063d5e48caea6a025d707a74f693974032b90af20af7686"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "466375c01a216c4ea1d90cc4ec5a1bde1e7fc555"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
