// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e8ea357d76fec0468ebf307f2abd3ff3d01b9b8e1d2bc2698a661a03a50dca84"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7b62a7cd184e7c96411f3fb78c41c7458c11d846"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
