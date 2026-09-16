// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e8ea357d76fec0468ebf307f2abd3ff3d01b9b8e1d2bc2698a661a03a50dca84"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1852f9488f605556fdd2e5d4f8c39d4ed3f0bbe7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
