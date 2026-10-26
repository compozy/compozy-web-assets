// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0c081da0f8e252b1b3fa0c0b5d72e72e1439bc414bd33612514ca610e16f72fb"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "65ade05b2f73585883d19b60a7fb9d5ec135a45a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
