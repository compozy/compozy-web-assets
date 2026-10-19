// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "123da120d1828997aeedb26bb581e19bce13642ec1414422692604dd04245e96"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8ea9bb5caf8be2c6287a714f679de5acfa2f86f8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
