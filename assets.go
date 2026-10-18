// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "42ba2c9e6f842eb5684db6b92c9f92275e4cac60060cdec0d0a4992e15f75975"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c624fab8216ae9fdd32d82ed547a03e1cec37781"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
