// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "786b1129cf503d8f458401fb187b1d50c96903f4670ffd5ed6b5b42691de0b2f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b880608dde5a9b93e6a8420eb09157ff569bd9eb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
