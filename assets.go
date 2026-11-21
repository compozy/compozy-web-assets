// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "786b1129cf503d8f458401fb187b1d50c96903f4670ffd5ed6b5b42691de0b2f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6f0354d098b1726794e7d25f2aa3d0a6536e1163"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
