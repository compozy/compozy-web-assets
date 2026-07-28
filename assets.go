// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e244a39c0c3a1eb651bf8958a55816134cf5b198df8af166f4db5feec105304a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "623f4736425b1a35c2b20c8d309064a81a40848a"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
