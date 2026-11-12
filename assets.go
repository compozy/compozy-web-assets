// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3341c8c0b09f2068ca4c2207f43ef0e817a9d3745d57385c2b9990519b0a8162"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "17e3f8b41e7638131a7a73f4626ac5efc195c7a7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
