// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1f7290731537e24845620f9582dd33cbad1f7a6aa6cb9cea3800674058912959"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "eb865c4f2f67b55bf310199dd7671a5a97a5d074"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
