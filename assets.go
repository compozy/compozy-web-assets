// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "06bd7080e36de5901e1cf86055d5f04203a7ebaf3368b1f258ddd4b2c88fd268"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "35cad0cddff7a55935ed0b77b2d3b5c5ba956662"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
