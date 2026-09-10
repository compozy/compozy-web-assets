// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "51f32bba9be3feb0260bfa02a896363b06f1fb2e63cea5a7053c4ddeba63d654"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "470c7c8f626be51436280457c2964f8a6093c2d2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
