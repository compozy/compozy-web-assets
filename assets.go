// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e244a39c0c3a1eb651bf8958a55816134cf5b198df8af166f4db5feec105304a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ddad19b84a5243a2dab18e087efbbee7d90221f7"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
