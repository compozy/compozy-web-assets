// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "51f32bba9be3feb0260bfa02a896363b06f1fb2e63cea5a7053c4ddeba63d654"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "af831a9f87aad9800f29d9f85bf4d1e5bd597614"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
