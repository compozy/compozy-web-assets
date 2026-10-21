// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e3a8fb2334f852d8fd79582f1b7014a6d05e862b560a0254a587f3a0d0dee929"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8234b2f3b7c87abb0bb9e55b098f400ae118bb67"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
