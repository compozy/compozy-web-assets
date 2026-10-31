// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d940a6840234c74bb1d6e26c04e4566bea14479a39ef2d48ccbf6fe14877fe87"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c9fedb65703b40fd86b25247263996c29b25e0cb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
