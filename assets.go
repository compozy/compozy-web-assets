// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f11bf0d7b9319fb26fb092e63c72b6993a6c7b1a2e68a16526abafd959d5f7ac"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4dea54017af74cba3901e24a416a4e4ae9c4d596"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
