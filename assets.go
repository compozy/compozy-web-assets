// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "223fd374e8e9c465c1f11f255ef4f4ba0baed1e009b53af8842e179361c3382c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6b7da9aa8c827ee9d3910306f37864bd1f4f23d8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
