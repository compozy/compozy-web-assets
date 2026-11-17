// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d6b44d8f40f484ee834edfd4dfef63270189691b9e75afc600056a73925fd442"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bed94f00410d66967c81a5b98ae924902726e887"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
