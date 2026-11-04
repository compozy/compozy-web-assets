// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "13dcebdc15f9d435b66c8b7b9c0b4f4c37e785866ab9876c62d7e732dc3ae8dc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3fcf51e9da0cd3f7cef6aea0bc2f13699261d6a6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
