// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a6c2f9daa47a2f7bb3b6193908f9a278d3bb6c787c01cb5f7353d887e5f81521"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4f947940f9915f47be24f5acbd4c0c20a93622cb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
