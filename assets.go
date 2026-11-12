// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3341c8c0b09f2068ca4c2207f43ef0e817a9d3745d57385c2b9990519b0a8162"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "dee9ffb3c9142bb02901ce106f34496f7418c5c2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
