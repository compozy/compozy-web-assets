// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ad5db0a8109bd2e06ba6e93be6a4bbfd561e2615af04392decc29fd387026638"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "14f57173e5cae483392cce90a06098fd04dcfda0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
