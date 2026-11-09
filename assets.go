// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "61fcb2c90f79ec9872dd587f28f8263c2759b2160dada5dfb1b17690296c5b79"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e6ecd1fccfd216e0f15b7e54c6aaad8b3bdb9c3b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
