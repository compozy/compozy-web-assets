// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e2b1acfe6dc9961f31aca460fd9832d9dd0ebd78231956e92813e1923f837a21"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7702673a8f55a1a40e0b64a0d3a16018f6a91e59"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
