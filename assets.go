// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c2ad9b0207d4e8b47be50197f5ad50447af4163c831f2261f0f8f03c9a75a1d7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e596826b215d193f5c1e6c88ce3b42deedfb8a23"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
