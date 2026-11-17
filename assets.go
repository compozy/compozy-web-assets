// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d6b44d8f40f484ee834edfd4dfef63270189691b9e75afc600056a73925fd442"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f7498334f33d94695c762071c6fd3ca02f29a120"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
