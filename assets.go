// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4ecc7bec7ca8ee3b5c8d2b051cc3d836b6601835c9db85a7099a0968b8a8251b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "48f6b14555643963cd4b10cf09d8705230ce0dd6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
