// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "62e0a3d308351bcf8d775a49855beff9e52bff1e8178caac3fe29d9860f3b773"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "94c2c199eb48b6baad7af2a372c5e56076134cad"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
