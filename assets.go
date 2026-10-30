// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7a69a071c6409848ebd676c04d46a8d1a329d258c245080ab349da8d9a18087c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "dcb2570f145e1ae52fe9d77464020b1bbc49c8b1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
