// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dd01a63c13e9102237e2ea6f6b23ae84e927bbb4552cff27b227f19e85e37cc9"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4cb2853b4302d4d9b361fe80b64c60892eb5e6ea"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
