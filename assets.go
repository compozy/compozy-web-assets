// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "61fcb2c90f79ec9872dd587f28f8263c2759b2160dada5dfb1b17690296c5b79"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a751dcac91af2cd4e91cc8d9c60bbbe66477863f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
