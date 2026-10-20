// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "37c1a037e93c73f51688085468cad7ede468afd358798d7460be3e38d4d235d8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "93c2ed58896e37f3f754a14d207441477f184992"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
