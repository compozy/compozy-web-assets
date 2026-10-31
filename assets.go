// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "878a4452b39ed5abe8357c034c3f78b4cdee976326e73ba7e6f4c20da6effcba"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2d6577c0e7d66d0fdc56a9e8f5209a28a375e5ae"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
