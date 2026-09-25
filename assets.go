// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "068796d9834ce4a82400bf1a134b0233e04ee8d0603e8e7f794ceb597a3513e4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bc8cb2a2b31c288e80b46cfb1fc8b2136a3c9de5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
