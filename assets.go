// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "878a4452b39ed5abe8357c034c3f78b4cdee976326e73ba7e6f4c20da6effcba"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b0f3612d0308cbc8b7859c6b8197dda93d906dcd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
