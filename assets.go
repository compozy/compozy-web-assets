// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1c170ef60e5bc2cbe184ecda78e2b87082e7f38d3fa0b1fe3d6aff2aebfa5e72"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "37d7adc0957e5f53915b7d9d174aa556d8662bdd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
