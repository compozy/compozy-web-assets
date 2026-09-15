// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "27f592fe3e3f082679d975f0625bc4bd48dc62e090adaf1590b509b202d983d7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a0c80e810459b2638d9ab3557b90efebc28a5f32"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
