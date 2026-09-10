// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "69c23f598dcada025c04f12d89cea8d353553269a816247da6bb5b1e7b970c26"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "29f4bee46fb0782cb6ae53f54fdaa0c68550f83c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
