// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bfbbaf6c8c2b1b4fddb404c3fd19a3308e3d5fae979ff717837adebe40fb4546"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f0036344c0a33a17a8e19202b5826444f62604e4"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
