// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8fbccf0915bf02ad0bc0f0fd430e87e0b9fcc5b3ab5a858ea616a2a5bded3968"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cb2032d651001f67a480af34f423787cf39d9eac"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
