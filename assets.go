// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "51f32bba9be3feb0260bfa02a896363b06f1fb2e63cea5a7053c4ddeba63d654"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6f51528922aad95187b6b3b267f101da661cb918"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
