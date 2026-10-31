// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d940a6840234c74bb1d6e26c04e4566bea14479a39ef2d48ccbf6fe14877fe87"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9e51edb16e35e23acd889775e6119a8916d2559f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
