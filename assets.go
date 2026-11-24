// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c890bfdf351ea97a5f2045aa0e7b3722d8bfce5e6bb3f32998f8aa2000f239ba"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9aa4d3866429af9695a8a65fb0f0e46f949e517b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
