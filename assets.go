// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b6b6bb9a6df999042b70b76388ab7f437cb05dc195b7067ceae2fe0e1cfdb26e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4dfb5f965f24b21c4aa8978270fbbd95ef46e4f8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
