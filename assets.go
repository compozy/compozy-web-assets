// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1eda069b4924948c30f9a82cdcaf3ac6b92fe0f0ec4998b721a6722c302b30a4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ca1a2dad3f084788fa8b8560559cdf44bd007e1b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
