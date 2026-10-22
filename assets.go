// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2bea4928d83b3b1e711057a50afb806e0a5bbc848f67f5539fdc68e335d3f615"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9798dbec859e7e6a00d46b0eb3ab6cae9dddf2d3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
