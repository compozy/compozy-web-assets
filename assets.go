// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "06039351747af74a981f3b92ce7f6c43f3679c593fc4365599e5269f94c66612"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7420b843bd43968b928946b69fd1273594cf969a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
