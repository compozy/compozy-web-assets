// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a23e6f1009e1e20427ddd272eff70b94a643e3d00aa7c21c947713a146614611"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "967fdae3ba4b2c72a1e443187cc235e4abed1f04"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
