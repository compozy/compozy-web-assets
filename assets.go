// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "37c1a037e93c73f51688085468cad7ede468afd358798d7460be3e38d4d235d8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2e63bba66c32ee210ac54902b07b5fca61536c5c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
