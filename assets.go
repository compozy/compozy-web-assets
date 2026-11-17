// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d6b44d8f40f484ee834edfd4dfef63270189691b9e75afc600056a73925fd442"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "802fc1c50c2d1213601676860f35bc9291b8f9f1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
