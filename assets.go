// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ad5db0a8109bd2e06ba6e93be6a4bbfd561e2615af04392decc29fd387026638"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "78d3a7f40d86d9026da69bde108b30bc28d1fac3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
