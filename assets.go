// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8fbccf0915bf02ad0bc0f0fd430e87e0b9fcc5b3ab5a858ea616a2a5bded3968"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "27b1323b8b6c4b3d5b31600c4007f02b102f1ab1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
