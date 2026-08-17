// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "859c990ef871144a60ca269acfb11819fa9b43fd364d7e41356048b010678b99"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5c5789e35f6d1a946d76231cba2329c88990ced9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
